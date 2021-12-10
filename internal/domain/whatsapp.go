package domain

import (
	"encoding/json"
	"mime/multipart"
)

type WhatsappSendTextForm struct {
	Msisdn      string `json:"msisdn" validate:"required"`
	Text        string `json:"text" validate:"required"`
	MsgQuotedID string `json:"msg_quoted_id"`
	MsgQuoted   string `json:"msg_quoted"`
}

type WhatsappSendLocationForm struct {
	Msisdn      string  `json:"msisdn" validate:"required"`
	Latitude    float64 `json:"latitude" validate:"required,latitude"`
	Longitude   float64 `json:"longitude" validate:"required,longitude"`
	MsgQuotedID string  `json:"msg_quoted_id"`
	MsgQuoted   string  `json:"msg_quoted"`
}

type WhatsappSendFileForm struct {
	Msisdn      string `json:"msisdn" validate:"required"`
	MsgQuotedID string `json:"msg_quoted_id"`
	MsgQuoted   string `json:"msg_quoted"`
	Message     string `json:"message"`
	File        string `json:"file" validate:"required,file"`
	FileHeader  *multipart.FileHeader
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
