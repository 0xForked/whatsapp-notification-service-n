package apperrors

import (
	"errors"
)

var (
	ErrAlreadyConnectedAndLoggedIn = errors.New("already connected & logged in")
	ErrInvalidSession              = errors.New("invalid session, service wasn't connected with your whatsapp")
	ErrPhoneNotConnected           = errors.New("something when wrong while trying to ping, please check phone connectivity")
	ErrInvalidFileFormat           = errors.New("invalid format, please try again. Accepted File (document, image, audio)")
	ErrRouteNotFound               = errors.New("http route not found")
	ErrMethodNotFound              = errors.New("http method not found")
)
