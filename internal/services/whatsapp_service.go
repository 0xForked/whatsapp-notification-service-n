package services

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"github.com/aasumitro/gowa/internal/domain/models"
	"github.com/aasumitro/gowa/internal/utils"
	"github.com/aasumitro/gowa/pkg/apperrors"
	"log"
	"os"
)

type WhatsappService struct {
	Conn *whatsapp.Conn
}

func (w *WhatsappService) Login() (qr string, err error) {
	if w.Conn.GetConnected() && w.Conn.GetLoggedIn() {
		err = apperrors.ErrAlreadyConnectedAndLoggedIn
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
			log.Println("ERROR", "error during login:", err)
		} else {
			log.Println("INFO", "login successful, session:", session)

			//save session
			if err := utils.WriteSession(session); err != nil {
				log.Println("ERROR", "error during save session:", err)
			}
		}

		return
	}()

	select {
	case qr = <-qrCodeChan:
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
	}

	return nil
}

func (w *WhatsappService) HasSession() (err error) {
	if w.Conn.Info == nil {
		_, _ = w.Conn.Disconnect()
	}

	if w.Conn.GetConnected() == false ||
		w.Conn.GetLoggedIn() == false {
		return apperrors.ErrInvalidSession
	}

	return nil
}

func (w *WhatsappService) Profile() (data map[string]string, err error) {
	conn := w.Conn
	ok, err := conn.AdminTest()
	if !ok {
		if err != nil {
			err = apperrors.ErrPhoneNotConnected
		}

		return
	}

	data = map[string]string{
		"display_name":   conn.Info.Pushname,
		"phone_battery":  fmt.Sprintf("%d%%", conn.Info.Battery),
		"phone_platform": conn.Info.Platform,
	}

	return
}

func (w *WhatsappService) SendText(form models.WhatsappSendTextForm) (msgId string, err error) {
	jid := utils.ParseMsisdn(form.Msisdn)

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: jid,
		},
		Text: form.Text,
	}

	msgId, err = w.Conn.Send(msg)

	return
}

func (w *WhatsappService) SendLocation(form models.WhatsappSendLocationForm) (msgId string, err error) {
	jid := utils.ParseMsisdn(form.Msisdn)

	msg := whatsapp.LocationMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: jid,
		},
		DegreesLatitude:  form.Latitude,
		DegreesLongitude: form.Longitude,
	}

	msgId, err = w.Conn.Send(msg)

	return
}

func (w *WhatsappService) SendFile(form models.WhatsappSendFileForm, fileType string) (msgId string, err error) {
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

	err = apperrors.ErrInvalidFileFormat
	return
}

func (w *WhatsappService) Logout() (err error) {
	if err := w.Conn.Logout(); err != nil {
		return err
	}

	if err := utils.RemoveSession(w.Conn); err != nil {
		return err
	}

	return
}

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

	msgId, err = w.Conn.Send(msg)

	return
}
