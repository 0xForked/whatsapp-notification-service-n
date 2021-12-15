package services

import (
	"github.com/Rhymen/go-whatsapp"
	"github.com/aasumitro/gowa/internal/domain"
	"github.com/aasumitro/gowa/internal/domain/models"
	"github.com/aasumitro/gowa/internal/utils"
	"os"
)

type WhatsappService struct {
	Conn *whatsapp.Conn
}

func (w *WhatsappService) Login() (qr string, err error) {
	if w.Conn.GetConnected() && w.Conn.GetLoggedIn() {
		err = domain.ErrAlreadyConnectedAndLoggedIn
		return
	}

	qrCode := make(chan string)
	qrCodeChan := make(chan string)
	go func() {
		qrStr := <-qrCode
		qrCodeChan <- qrStr
	}()

	go func() {
		session := whatsapp.Session{}
		session, err = w.Conn.Login(qrCode)
		if err != nil {
			// log.Println(log.LogLevelError, "error during login:", err)
		} else {
			// log.Println(log.LogLevelInfo, "login successful, session:", session)

			//save session
			if err := utils.WriteSession(session); err != nil {
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

func (w *WhatsappService) RestoreSession() error {
	//load saved session
	session, err := utils.ReadSession()
	if err == nil {
		//restore session
		session, err = w.Conn.RestoreWithSession(session)
		if err != nil {
			_ = os.Remove(os.Getenv("WAC_SESSION_PATH") +
				"/whatsapp_session.gob")
			return err
		}

		//save session
		err = utils.WriteSession(session)
		if err != nil {
			return err
		}

		// w.whatsappConn.AddHandler(utils.WhatsappHandler{})
	}

	return nil
}

func (w *WhatsappService) HasSession() (err error) {
	if w.Conn.GetConnected() == false ||
		w.Conn.GetLoggedIn() == false {
		return domain.ErrInvalidSession
	}

	return nil
}

func (w *WhatsappService) SendText(
	form models.WhatsappSendTextForm,
) (msgId string, err error) {
	jid := utils.ParseMsisdn(form.Msisdn)

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: jid,
		},
		Text: form.Text,
	}

	//if len(form.MsgQuotedID) != 0 {
	//	quotedMessage := proto.Message{
	//		Conversation: &form.MsgQuoted,
	//	}
	//
	//	ContextInfo := whatsapp.ContextInfo{
	//		QuotedMessageID: form.MsgQuotedID,
	//		QuotedMessage:   &quotedMessage,
	//		Participant:     jid,
	//	}
	//
	//	msg.ContextInfo = ContextInfo
	//}

	msgId, err = w.Conn.Send(msg)

	return
}

func (w *WhatsappService) SendLocation(
	form models.WhatsappSendLocationForm,
) (msgId string, err error) {
	jid := utils.ParseMsisdn(form.Msisdn)

	msg := whatsapp.LocationMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: jid,
		},
		DegreesLatitude:  form.Latitude,
		DegreesLongitude: form.Longitude,
	}

	//if len(form.MsgQuotedID) != 0 {
	//	quotedMessage := proto.Message{
	//		Conversation: &form.MsgQuoted,
	//	}
	//
	//	ContextInfo := whatsapp.ContextInfo{
	//		QuotedMessageID: form.MsgQuotedID,
	//		QuotedMessage:   &quotedMessage,
	//		Participant:     jid,
	//	}
	//
	//	msg.ContextInfo = ContextInfo
	//}

	msgId, err = w.Conn.Send(msg)

	return
}

func (w *WhatsappService) SendFile(
	form models.WhatsappSendFileForm,
	fileType string,
) (msgId string, err error) {
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

func (w *WhatsappService) Logout() (err error) {
	err = utils.LogoutSession(w.Conn)
	if err != nil {
		return
	}

	return
}

//func NewWhatsappService(conn *whatsapp.Conn) domain.WhatsappServiceContract {
//	return &whatsappService{whatsappConn: conn}
//}

func sendDocument(w *WhatsappService, form models.WhatsappSendFileForm) (msgId string, err error) {
	jid := utils.ParseMsisdn(form.Msisdn)

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

	//if len(form.MsgQuotedID) != 0 {
	//	quotedMessage := proto.Message{
	//		Conversation: &form.MsgQuoted,
	//	}
	//
	//	ContextInfo := whatsapp.ContextInfo{
	//		QuotedMessageID: form.MsgQuotedID,
	//		QuotedMessage:   &quotedMessage,
	//		Participant:     jid,
	//	}
	//
	//	msg.ContextInfo = ContextInfo
	//}

	msgId, err = w.Conn.Send(msg)

	return
}

func sendImage(w *WhatsappService, form models.WhatsappSendFileForm) (msgId string, err error) {
	jid := utils.ParseMsisdn(form.Msisdn)

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

	//if len(form.MsgQuotedID) != 0 {
	//	quotedMessage := proto.Message{
	//		Conversation: &form.MsgQuoted,
	//	}
	//
	//	ContextInfo := whatsapp.ContextInfo{
	//		QuotedMessageID: form.MsgQuotedID,
	//		QuotedMessage:   &quotedMessage,
	//		Participant:     jid,
	//	}
	//
	//	msg.ContextInfo = ContextInfo
	//}

	msgId, err = w.Conn.Send(msg)

	return
}

func sendAudio(w *WhatsappService, form models.WhatsappSendFileForm) (msgId string, err error) {
	jid := utils.ParseMsisdn(form.Msisdn)

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

	//if len(form.MsgQuotedID) != 0 {
	//	quotedMessage := proto.Message{
	//		Conversation: &form.MsgQuoted,
	//	}
	//
	//	ContextInfo := whatsapp.ContextInfo{
	//		QuotedMessageID: form.MsgQuotedID,
	//		QuotedMessage:   &quotedMessage,
	//		Participant:     jid,
	//	}
	//
	//	msg.ContextInfo = ContextInfo
	//}

	msgId, err = w.Conn.Send(msg)

	return
}

// check if phone is connected to internet
func isPhoneConnected(w *WhatsappService) error {
	conn := w.Conn
	ok, err := conn.AdminTest()
	if !ok {
		if err != nil {
			return err
		}

		return domain.ErrPhoneNotConnected
	}

	return nil
}
