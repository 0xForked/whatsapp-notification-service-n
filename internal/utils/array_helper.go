package utils

import "strings"

func InArray(needle interface{}, haystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range haystack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range haystack.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range haystack.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}

	return false
}

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}
