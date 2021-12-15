package domain

import (
	"encoding/json"
	"mime/multipart"
)

var RequestWhatsappErrorMessage = map[string]string{
	"error_required_msisdn":    "msisdn (mobile subscriber integrated services digital network number) is required",
	"error_required_text":      "text message is required",
	"error_required_latitude":  "latitude is required",
	"error_required_longitude": "longitude is required",
	"error_required_file":      "file is required",
}

type WhatsappSendTextForm struct {
	Msisdn string `json:"msisdn" form:"msisdn" validate:"required" binding:"required" msg:"error_required_msisdn"`
	Text   string `json:"text" form:"text" validate:"required" binding:"required" msg:"error_required_text"`
	//MsgQuotedID string `json:"msg_quoted_id" form:"msg_quoted_id"`
	//MsgQuoted   string `json:"msg_quoted" form:"msg_quoted"`
}

type WhatsappSendLocationForm struct {
	Msisdn    string  `json:"msisdn" form:"msisdn" validate:"required" binding:"required" msg:"error_required_msisdn"`
	Latitude  float64 `json:"latitude" form:"latitude" validate:"required,latitude" binding:"required" msg:"error_required_latitude"`
	Longitude float64 `json:"longitude" form:"longitude" validate:"required,longitude" binding:"required" msg:"error_required_longitude"`
	//MsgQuotedID string  `json:"msg_quoted_id" form:"msg_quoted_id"`
	//MsgQuoted   string  `json:"msg_quoted" form:"msg_quoted"`
}

type WhatsappSendFileForm struct {
	Msisdn     string                `json:"msisdn" form:"msisdn" validate:"required" binding:"required" msg:"error_required_msisdn"`
	Message    string                `json:"message" form:"message"`
	FileHeader *multipart.FileHeader `json:"file" form:"file" binding:"required" msg:"error_required_file"`
	//MsgQuotedID string                `json:"msg_quoted_id" form:"msg_quoted_id"`
	//MsgQuoted   string                `json:"msg_quoted" form:"msg_quoted"`
}

type WhatsappWebServer struct {
	Version struct {
		Major int
		Minor int
		Build int
	}
}

type WhatsappWebClient struct {
	Version struct {
		Major int
		Minor int
		Build int
	}
}

type WhatsappWeb struct {
	Server WhatsappWebServer
	Client WhatsappWebClient
}

type Whatsapp struct {
	WhatsappWeb   WhatsappWeb
	SessionJid    string
	SessionID     string
	SessionFile   string
	SessionStart  uint64
	ReconnectTime int
}

// FromJSON decode json to book struct
func (w *Whatsapp) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, w)
}

// ToJSON encode book struct to json
func (w *Whatsapp) ToJSON() []byte {
	str, _ := json.Marshal(w)
	return str
}
