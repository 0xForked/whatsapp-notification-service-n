package domain

type TextMessage struct {
	Msisdn string `json:"msisdn" form:"msisdn"`
	Text   string `json:"text" form:"text"`
}
