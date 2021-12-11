package domain

// WhatsappServiceContract represent the whatsapp's use cases
type WhatsappServiceContract interface {
	// Login Logout RestoreSession - for whatsapp account connection
	Login() (qr string, err error)
	Logout() (err error)
	RestoreSession() error
	// HasSession - check if the session is valid
	HasSession() error
	// SendText SendLocation SendFile - for whatsapp message
	SendText(form WhatsappSendTextForm) (msgId string, err error)
	SendLocation(form WhatsappSendLocationForm) (msgId string, err error)
	SendFile(form WhatsappSendFileForm, fileType string) (msgId string, err error)
}
