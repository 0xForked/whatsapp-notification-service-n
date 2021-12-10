package services

import (
	"encoding/gob"
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"github.com/Rhymen/go-whatsapp/binary/proto"
	"github.com/aasumitro/gowa/internal/domain"
	"os"
	"strings"
	"time"
)

type whatsappService struct {
	whatsappConn *whatsapp.Conn
}

func (w *whatsappService) Login(
	vMajor, vMinor, vBuild, timeout, reconnect int,
	clientNameShort, clientNameLong string,
) (qr string, err error) {
	if w.whatsappConn.GetConnected() && w.whatsappConn.GetLoggedIn() {
		err = domain.ErrAlreadyConnectedAndLoggedIn
		return
	}

	w.whatsappConn, err = whatsapp.NewConnWithOptions(&whatsapp.Options{
		Timeout:         time.Duration(timeout) * time.Second,
		ShortClientName: clientNameShort,
		LongClientName:  clientNameLong,
	})
	if err != nil {
		return
	}

	info, err := syncVersion(
		w.whatsappConn, vMajor, vMinor, vBuild,
	)
	if err != nil {
		return
	}
	fmt.Println(info)

	// log.Println(log.LogLevelInfo, "whatsapp-session-init", info)
	// w.whatsappConn.AddHandler(utils.WhatsappHandler{})

	qrCode := make(chan string)
	qrCodeChan := make(chan string)
	go func() {
		qrStr := <-qrCode
		qrCodeChan <- qrStr
	}()

	go func() {
		session := whatsapp.Session{}
		session, err = w.whatsappConn.Login(qrCode)
		if err != nil {
			// log.Println(log.LogLevelError, "error during login:", err)
		} else {
			// log.Println(log.LogLevelInfo, "login successful, session:", session)

			//save session
			if err := writeSession(session); err != nil {
				// log.Println(log.LogLevelError, "error during login:", err)
			}

			// w.whatsappConn.AddHandler(utils.WhatsappHandler{})
		}

		return
	}()

	select {
	case qr = <-qrCodeChan:
		// Test ping
		err = isPhoneConnected(w)
		if err != nil {
			// log.Println(log.LogLevelError, "error during login:", err)
			return
		}
		return
	}
}

func (w *whatsappService) RestoreSession() error {
	//load saved session
	session, err := readSession()
	if err == nil {
		//restore session
		session, err = w.whatsappConn.RestoreWithSession(session)
		if err != nil {
			_ = os.Remove(os.Getenv("WHATSAPP_CLIENT_SESSION_PATH") +
				"/whatsapp_session.gob")
			return err
		}

		//save session
		err = writeSession(session)
		if err != nil {
			return err
		}

		// w.whatsappConn.AddHandler(utils.WhatsappHandler{})
	}

	return nil
}

func (w *whatsappService) GetInfo() (info domain.WhatsappWeb, err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = domain.ErrInvalidSession
		return
	}

	v := w.whatsappConn.GetClientVersion()
	info.Client.Version.Major = v[0]
	info.Client.Version.Minor = v[1]
	info.Client.Version.Build = v[2]

	v, err = whatsapp.CheckCurrentServerVersion()
	info.Server.Version.Major = v[0]
	info.Server.Version.Minor = v[1]
	info.Server.Version.Build = v[2]
	if err != nil {
		panic(err)
	}

	return
}

func (w *whatsappService) SendText(form domain.WhatsappSendTextForm) (msgId string, err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = domain.ErrInvalidSession
		return
	}

	jid := parseMsisdn(form.Msisdn)

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: jid,
		},
		Text: form.Text,
	}

	if len(form.MsgQuotedID) != 0 {
		quotedMessage := proto.Message{
			Conversation: &form.MsgQuoted,
		}

		ContextInfo := whatsapp.ContextInfo{
			QuotedMessageID: form.MsgQuotedID,
			QuotedMessage:   &quotedMessage,
			Participant:     jid,
		}

		msg.ContextInfo = ContextInfo
	}

	msgId, err = w.whatsappConn.Send(msg)

	return
}

func (w *whatsappService) SendLocation(form domain.WhatsappSendLocationForm) (msgId string, err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = domain.ErrInvalidSession
		return
	}

	jid := parseMsisdn(form.Msisdn)

	msg := whatsapp.LocationMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: jid,
		},
		DegreesLatitude:  form.Latitude,
		DegreesLongitude: form.Longitude,
	}

	if len(form.MsgQuotedID) != 0 {
		quotedMessage := proto.Message{
			Conversation: &form.MsgQuoted,
		}

		ContextInfo := whatsapp.ContextInfo{
			QuotedMessageID: form.MsgQuotedID,
			QuotedMessage:   &quotedMessage,
			Participant:     jid,
		}

		msg.ContextInfo = ContextInfo
	}

	msgId, err = w.whatsappConn.Send(msg)

	return
}

func (w *whatsappService) SendFile(form domain.WhatsappSendFileForm, fileType string) (msgId string, err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = domain.ErrInvalidSession
		return
	}

	switch fileType {
	case "document":
		msgId, err = sendDocument(w, form)
		return
	case "image":
		msgId, err = sendImage(w, form)
		return
	case "audio":
		msgId, err = sendAudio(w, form)
		return
	}

	err = domain.ErrInvalidFileFormat
	return
}

func (w *whatsappService) Logout() (err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = domain.ErrInvalidSession
		return
	}

	err = logout(w.whatsappConn)
	if err != nil {
		return
	}

	return
}

func (w *whatsappService) Groups(jid string) (g string, err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = domain.ErrInvalidSession
		return
	}

	data, err := w.whatsappConn.GetGroupMetaData(parseMsisdn(jid))
	if err != nil {
		return
	}

	g = <-data

	return
}

func NewWhatsappService(conn *whatsapp.Conn) domain.WhatsappServiceContact {
	return &whatsappService{whatsappConn: conn}
}

// read session from temporary file
func readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Create(
		os.Getenv("WHATSAPP_CLIENT_SESSION_PATH") +
			"/whatsapp_session.gob")
	if err != nil {
		return session, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

// Write whatsapp session on temporary directory
func writeSession(session whatsapp.Session) error {
	file, err := os.Create(
		os.Getenv("WHATSAPP_CLIENT_SESSION_PATH") +
			"/whatsapp_session.gob")
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}

// check if phone is connected to internet
func isPhoneConnected(w *whatsappService) error {
	conn := w.whatsappConn
	ok, err := conn.AdminTest()
	if !ok {
		if err != nil {
			return err
		}

		return domain.ErrPhoneNotConnected
	}

	return nil
}

func parseMsisdn(msisdn string) string {
	components := strings.Split(msisdn, "@")

	if len(components) > 1 {
		msisdn = components[0]
	}

	suffix := "@s.whatsapp.net"

	if len(strings.SplitN(msisdn, "-", 2)) == 2 {
		suffix = "@g.us"
	}

	return msisdn + suffix
}

func syncVersion(
	conn *whatsapp.Conn,
	versionClientMajor int,
	versionClientMinor int,
	versionClientBuild int,
) (string, error) {
	conn.SetClientVersion(versionClientMajor, versionClientMinor, versionClientBuild)
	versionClient := conn.GetClientVersion()
	return fmt.Sprintf(
		"whatsapp version %v.%v.%v",
		versionClient[0],
		versionClient[1],
		versionClient[2],
	), nil
}

func sendDocument(w *whatsappService, form domain.WhatsappSendFileForm) (msgId string, err error) {
	jid := parseMsisdn(form.Msisdn)

	f, err := form.FileHeader.Open()
	if err != nil {
		return "", err
	}

	msg := whatsapp.DocumentMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: jid,
		},
		Content:  f,
		Type:     form.FileHeader.Header.Get("Content-Type"),
		FileName: form.FileHeader.Filename,
		Title:    form.FileHeader.Filename,
	}

	if len(form.MsgQuotedID) != 0 {
		quotedMessage := proto.Message{
			Conversation: &form.MsgQuoted,
		}

		ContextInfo := whatsapp.ContextInfo{
			QuotedMessageID: form.MsgQuotedID,
			QuotedMessage:   &quotedMessage,
			Participant:     jid,
		}

		msg.ContextInfo = ContextInfo
	}

	msgId, err = w.whatsappConn.Send(msg)

	return
}

func sendImage(w *whatsappService, form domain.WhatsappSendFileForm) (msgId string, err error) {
	jid := parseMsisdn(form.Msisdn)

	f, err := form.FileHeader.Open()
	if err != nil {
		return "", err
	}

	msg := whatsapp.ImageMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: jid,
		},
		Content: f,
		Type:    form.FileHeader.Header.Get("Content-Type"),
		Caption: form.Message,
	}

	if len(form.MsgQuotedID) != 0 {
		quotedMessage := proto.Message{
			Conversation: &form.MsgQuoted,
		}

		ContextInfo := whatsapp.ContextInfo{
			QuotedMessageID: form.MsgQuotedID,
			QuotedMessage:   &quotedMessage,
			Participant:     jid,
		}

		msg.ContextInfo = ContextInfo
	}

	msgId, err = w.whatsappConn.Send(msg)

	return
}

func sendAudio(w *whatsappService, form domain.WhatsappSendFileForm) (msgId string, err error) {
	jid := parseMsisdn(form.Msisdn)

	f, err := form.FileHeader.Open()
	if err != nil {
		return "", err
	}

	msg := whatsapp.AudioMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: jid,
		},
		Content: f,
		Type:    form.FileHeader.Header.Get("Content-Type"),
	}

	if len(form.MsgQuotedID) != 0 {
		quotedMessage := proto.Message{
			Conversation: &form.MsgQuoted,
		}

		ContextInfo := whatsapp.ContextInfo{
			QuotedMessageID: form.MsgQuotedID,
			QuotedMessage:   &quotedMessage,
			Participant:     jid,
		}

		msg.ContextInfo = ContextInfo
	}

	msgId, err = w.whatsappConn.Send(msg)

	return
}

func logout(wac *whatsapp.Conn) error {
	defer func() {
		fmt.Println("Disconnecting..")
		_, _ = wac.Disconnect()
	}()

	err := wac.Logout()
	if err != nil {
		return err
	}

	_ = os.Remove(os.Getenv("WHATSAPP_CLIENT_SESSION_PATH") +
		"/whatsapp_session.gob")

	fmt.Println("Logout success..")

	return nil
}
