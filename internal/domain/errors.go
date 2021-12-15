package domain

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

var (
	ErrAlreadyConnectedAndLoggedIn = errors.New("already connected & logged in")
	ErrInvalidSession              = errors.New("invalid session, service wasn't connected with your whatsapp")
	ErrLoginInProgress             = errors.New("login or restore already running")
	ErrNotConnected                = errors.New("not connected")
	ErrInvalidWsData               = errors.New("received invalid data")
	ErrInvalidWsState              = errors.New("can't handle binary data when not logged in")
	ErrConnectionTimeout           = errors.New("connection timed out")
	ErrMissingMessageTag           = errors.New("no messageTag specified or to short")
	ErrInvalidHmac                 = errors.New("invalid hmac")
	ErrInvalidServerResponse       = errors.New("invalid response received from server")
	ErrServerRespondedWith404      = errors.New("server responded with status 404")
	ErrMediaDownloadFailedWith404  = errors.New("download failed with status code 404")
	ErrMediaDownloadFailedWith410  = errors.New("download failed with status code 410")
	ErrInvalidWebsocket            = errors.New("invalid websocket")
	ErrPhoneNotConnected           = errors.New("something when wrong while trying to ping, please check phone connectivity")
	ErrOptionsNotProvided          = errors.New("new conn options not provided")
	ErrInvalidFileFormat           = errors.New("invalid format, please try again. Accepted File (document, image, audio)")

	ErrRouteNotFound  = errors.New("http route not found")
	ErrMethodNotFound = errors.New("http method not found")
)

type GinErrors interface {
	ListAllErrors(model interface{}, err error) map[string]string
}

type ginErrors struct {
	errorMaps map[string]string
}

type ErrorResult struct {
	Field   string
	JsonTag string
	Message string
}

func NewGinErrors(errors map[string]string) GinErrors {
	return &ginErrors{
		errorMaps: errors,
	}
}

func (ge ginErrors) ListAllErrors(model interface{}, err error) map[string]string {
	errs := map[string]string{}
	fields := map[string]ErrorResult{}

	if _, ok := err.(validator.ValidationErrors); ok {
		// resolve all json tags for the struct
		types := reflect.TypeOf(model)
		values := reflect.ValueOf(model)

		for i := 0; i < types.NumField(); i++ {
			field := types.Field(i)
			value := values.Field(i).Interface()
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = field.Name
			}
			messageTag := field.Tag.Get("msg")
			msg := ge.getErrorMessage(messageTag)

			fmt.Printf("%s: %v = %v, tag= %v\n", field.Name, field.Type, value, jsonTag)
			fields[field.Name] = ErrorResult{
				Field:   field.Name,
				JsonTag: jsonTag,
				Message: msg,
			}
		}

		for _, e := range err.(validator.ValidationErrors) {
			if field, ok := fields[e.Field()]; ok {
				if field.Message != "" {
					errs[field.JsonTag] = field.Message
				} else {
					errs[field.JsonTag] = e.Error()
				}
			}
		}
	} else {
		errs["0"] = err.Error()
	}

	return errs
}

func (ge ginErrors) getErrorMessage(key string) string {
	if value, ok := ge.errorMaps[key]; ok {
		return value
	}
	return key
}
