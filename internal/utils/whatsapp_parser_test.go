package utils_test

import (
	"github.com/aasumitro/gowa/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMsisdn(t *testing.T) {
	var testsData = []struct {
		input    string
		expected string
	}{
		{"+1234", "+1234@s.whatsapp.net"},
		{"+8271112233", "+8271112233@s.whatsapp.net"},
		{"+6282271112233", "+6282271112233@s.whatsapp.net"},
		{"12345@msisdn", "12345@s.whatsapp.net"},
		{"+12345@msisdn", "+12345@s.whatsapp.net"},
		{"-12345", "-12345@g.us"},
		{"-12345@msisdn", "-12345@g.us"},
	}

	for _, test := range testsData {
		actual := utils.ParseMsisdn(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
