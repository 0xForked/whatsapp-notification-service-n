package domain

// WhatsappService represent the whatsapp's use cases
type WhatsappService interface {
	RestoreSession() error
	GetInfo() (info WhatsappWeb, err error)
	SendText(form WhatsappSendTextForm) (msgId string, err error)
	SendLocation(form WhatsappSendLocationForm) (msgId string, err error)
	SendFile(form WhatsappSendFileForm, fileType string) (msgId string, err error)
	Logout() (err error)
	Groups(jid string) (g string, err error)
}

//Login(vMajor, vMinor, vBuild, timeout, reconnect int, clientNameShort, clientNameLong string) (qrCode string, err error)
