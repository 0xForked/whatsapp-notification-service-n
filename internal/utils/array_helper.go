package utils

import (
	"reflect"
	"strings"
)

func InArray[T any](needle T, haystack []T) bool {
	for _, item := range haystack {
		if reflect.DeepEqual(needle, item) {
			return true
		}
	}

	return false
}

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	}

	return strings.Split(text, delimiter)
}
