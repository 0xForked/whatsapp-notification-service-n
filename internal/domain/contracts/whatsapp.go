package contracts

import "github.com/aasumitro/gowa/internal/domain/models"

// WhatsappService represent the whatsapp's use cases
type WhatsappService interface {
	// Login Logout RestoreSession - for whatsapp account connection
	Login() (qr string, err error)
	Logout() (err error)
	RestoreSession() error
	// HasSession - check if the session is valid
	HasSession() error
	Profile() (data map[string]string, err error)
	// SendText SendLocation SendFile - for whatsapp message
	SendText(form models.WhatsappSendTextForm) (msgId string, err error)
	SendLocation(form models.WhatsappSendLocationForm) (msgId string, err error)
	SendFile(form models.WhatsappSendFileForm, fileType string) (msgId string, err error)
}
