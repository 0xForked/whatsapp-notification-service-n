package domain

import "mime/multipart"

type MessageType int
type FileType int

const (
	MessageTypeText MessageType = iota
	MessageTypeLocation
	MessageTypeFile
)

const (
	FileTypeDocument FileType = iota
	FileTypeImage
	FileTypeAudio
)

type (
	TextMessage struct {
		Msisdn string
		Text   string
	}

	LocationMessage struct {
		Msisdn    string
		Latitude  float64
		Longitude float64
	}

	FileMessage struct {
		Msisdn     string
		Message    string
		FileHeader *multipart.FileHeader
		Type       FileType
	}
)
