package domain

// WhatsappServiceContact WhatsappService represent the whatsapp's use cases
type WhatsappServiceContact interface {
	Login(vMajor, vMinor, vBuild, timeout, reconnect int, clientNameShort, clientNameLong string) (qr string, err error)
	RestoreSession() error
	GetInfo() (info WhatsappWeb, err error)
	SendText(form WhatsappSendTextForm) (msgId string, err error)
	SendLocation(form WhatsappSendLocationForm) (msgId string, err error)
	SendFile(form WhatsappSendFileForm, fileType string) (msgId string, err error)
	Logout() (err error)
	Groups(jid string) (g string, err error)
}
