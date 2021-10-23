package service

import (
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"github.com/Rhymen/go-whatsapp/binary/proto"
	"github.com/aasumitro/wagorf/domain"
	"github.com/aasumitro/wagorf/utils"
	"github.com/aasumitro/wagorf/utils/log"
	"os"
	"strings"
	"time"
)

type whatsAppService struct {
	whatsappConn *whatsapp.Conn
}

func NewWhatsAppService(conn *whatsapp.Conn) domain.WhatsappService {
	return &whatsAppService{whatsappConn: conn}
}

func (w *whatsAppService) Login(vMajor, vMinor, vBuild, timeout, reconnect int, clientNameShort, clientNameLong string) (qrCodeStr string, err error) {
	if w.whatsappConn.GetConnected() && w.whatsappConn.GetLoggedIn() {
		err = errors.New("session already active")
		return
	}

	w.whatsappConn, err = whatsapp.NewConnWithOptions(&whatsapp.Options{
		// timeout
		Timeout: time.Duration(timeout) * time.Second,
		//Proxy:   proxy,
		// set custom client name
		ShortClientName: clientNameShort,
		LongClientName:  clientNameLong,
	})
	if err != nil {
		return
	}

	info, err := syncVersion(w.whatsappConn, vMajor, vMinor, vBuild)
	if err != nil {
		return
	}
	log.Println(log.LogLevelInfo, "whatsapp-session-init", info)

	w.whatsappConn.AddHandler(utils.WhatsappHandler{})

	qr := make(chan string)
	qrCodeChan := make(chan string)
	go func() {
		qrStr := <-qr
		qrCodeChan <- qrStr
	}()

	go func() {
		session := whatsapp.Session{}
		session, err = w.whatsappConn.Login(qr)
		if err != nil {
			log.Println(log.LogLevelError, "error during login:", err)
		} else {
			log.Println(log.LogLevelInfo, "login successful, session:", session)

			//save session
			err = writeSession(session)
			if err != nil {
				log.Println(log.LogLevelError, "error during login:", err)
			}

			w.whatsappConn.AddHandler(utils.WhatsappHandler{})
		}

		return
	}()

	select {
	case qrCodeStr = <-qrCodeChan:
		// Test ping
		/*err = testPing(w)
		if err != nil {
			log.Println(log.LogLevelError, "error during login:", err)
			return
		}*/
		return
	}
}

func (w *whatsAppService) GetInfo() (info domain.WaWeb, err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = errors.New("invalid session, please login")
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

func (w *whatsAppService) SendText(form domain.WaSendTextForm) (msgId string, err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = errors.New("invalid session, please login")
		return
	}

	jid := parseMsisdn(form.Msisdn)

	//072217ED965D0C89DC6A

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

func (w *whatsAppService) SendLocation(form domain.WaSendLocationForm) (msgId string, err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = errors.New("invalid session, please login")
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

func (w *whatsAppService) SendFile(form domain.WaSendFileForm, fileType string) (msgId string, err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = errors.New("invalid session, please login")
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
	case "video":
		msgId, err = sendVideo(w, form)
		return
	}

	err = errors.New("invalid format, please try again")
	return
}

func (w *whatsAppService) Groups(jid string) (g string, err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = errors.New("invalid session, please login")
		return
	}

	data, err := w.whatsappConn.GetGroupMetaData(parseMsisdn(jid))
	if err != nil {
		return
	}

	g = <-data

	return
}

func (w *whatsAppService) Logout() (err error) {
	if w.whatsappConn.GetConnected() == false || w.whatsappConn.GetLoggedIn() == false {
		err = errors.New("invalid session, please login")
		return
	}

	err = logout(w.whatsappConn)
	if err != nil {
		return
	}

	return
}

func (w *whatsAppService) RestoreSession() error {
	//load saved session
	session, err := readSession()
	if err == nil {
		//restore session
		session, err = w.whatsappConn.RestoreWithSession(session)
		if err != nil {
			_ = os.Remove(os.Getenv("WHATSAPP_CLIENT_SESSION_PATH") + "/whatsappSession.gob")
			return err
		}

		//save session
		err = writeSession(session)
		if err != nil {
			return err
		}

		w.whatsappConn.AddHandler(utils.WhatsappHandler{})
	}

	return nil
}

func readSession() (whatsapp.Session, error) {
	fmt.Println(os.TempDir())

	session := whatsapp.Session{}
	//file, err := os.Open(os.TempDir() + "/whatsappSession.gob")
	file, err := os.Open(os.Getenv("WHATSAPP_CLIENT_SESSION_PATH") + "/whatsappSession.gob")
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

func writeSession(session whatsapp.Session) error {
	fmt.Println(os.TempDir())
	//file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	file, err := os.Create(os.Getenv("WHATSAPP_CLIENT_SESSION_PATH") + "/whatsappSession.gob")
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

func testPing(w *whatsAppService) error {
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

func sendImage(w *whatsAppService, form domain.WaSendFileForm) (msgId string, err error) {
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

func sendAudio(w *whatsAppService, form domain.WaSendFileForm) (msgId string, err error) {
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

func sendVideo(w *whatsAppService, form domain.WaSendFileForm) (msgId string, err error) {
	jid := parseMsisdn(form.Msisdn)

	f, err := form.FileHeader.Open()
	if err != nil {
		return "", err
	}

	msg := whatsapp.VideoMessage{
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

func sendDocument(w *whatsAppService, form domain.WaSendFileForm) (msgId string, err error) {
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

func logout(wac *whatsapp.Conn) error {
	defer func() {
		fmt.Println("Disconnecting..")
		_, _ = wac.Disconnect()
	}()

	err := wac.Logout()
	if err != nil {
		return err
	}

	_ = os.Remove(os.TempDir() + "/whatsappSession.gob")

	fmt.Println("Logout success..")

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

func syncVersion(conn *whatsapp.Conn, versionClientMajor int, versionClientMinor int, versionClientBuild int) (string, error) {
	// Bug Happend When Using This Function
	// Then Set Manualy WhatsApp Client Version
	// versionServer, err := whatsapp.CheckCurrentServerVersion()
	// if err != nil {
	// 	return "", err
	// }

	conn.SetClientVersion(versionClientMajor, versionClientMinor, versionClientBuild)
	versionClient := conn.GetClientVersion()

	return fmt.Sprintf("whatsapp version %v.%v.%v", versionClient[0], versionClient[1], versionClient[2]), nil
}
