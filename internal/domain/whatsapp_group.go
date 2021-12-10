package domain

import "encoding/json"

type WhatsAppGroup struct {
	ID           string          `json:"id"`
	Owner        string          `json:"owner"`
	Subject      string          `json:"subject"`
	Creation     int             `json:"creation"`
	SubjectTime  int             `json:"subjectTime"`
	SubjectOwner string          `json:"subjectOwner"`
	Desc         string          `json:"desc"`
	DescId       string          `json:"descId"`
	DescTime     int             `json:"descTime"`
	DescOwner    string          `json:"descOwner"`
	Participants []WhatsAppGroup `json:"participants"`
}

type WhatsAppGroupGroupParticipants struct {
	ID           string `json:"id"`
	IsAdmin      bool   `json:"isAdmin"`
	IsSuperAdmin bool   `json:"isSuperAdmin"`
}

// FromJSON decode json to book struct
func (w *WhatsAppGroup) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, w)
}

// ToJSON encode book struct to json
func (w *WhatsAppGroup) ToJSON() []byte {
	str, _ := json.Marshal(w)
	return str
}
