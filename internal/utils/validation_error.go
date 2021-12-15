package utils

import (
	"fmt"
	"github.com/aasumitro/gowa/internal/domain/contracts"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type validationErrors struct {
	errorMaps map[string]string
}

type ValidationErrorResult struct {
	Field   string
	JsonTag string
	Message string
}

func NewValidationErrors(errors map[string]string) contracts.ValidationErrors {
	return &validationErrors{
		errorMaps: errors,
	}
}

func (ge validationErrors) All(model interface{}, err error) map[string]string {
	errs := map[string]string{}
	fields := map[string]ValidationErrorResult{}

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
			fields[field.Name] = ValidationErrorResult{
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

func (ge validationErrors) getErrorMessage(key string) string {
	if value, ok := ge.errorMaps[key]; ok {
		return value
	}
	return key
}
