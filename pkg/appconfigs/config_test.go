package appconfigs_test

import (
	"github.com/aasumitro/gowa/pkg/appconfigs"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppConfig(t *testing.T) {
	viper.Reset()
	viper.SetConfigFile("../../.example.env")
	viper.SetConfigType("dotenv")
	appconfigs.LoadEnv()

	tt := []struct {
		name     string
		value    any
		expected any
		reflect  any
	}{
		{
			name:     "Test AppName Env",
			value:    appconfigs.Instance.AppName,
			expected: "Whatsapp Notification Service",
		},
		{
			name:     "Test AppVersion Env",
			value:    appconfigs.Instance.AppVersion,
			expected: "0.0.1-dev",
		},
		{
			name:     "Test AppUrl Env",
			value:    appconfigs.Instance.AppURL,
			expected: "localhost:8000",
		},
		{
			name:     "Test AppSessionPath Env",
			value:    appconfigs.Instance.AppSessionPath,
			expected: "./storage/sessions",
		},
		{
			name:     "Test AppUploadPath Env",
			value:    appconfigs.Instance.AppUploadPath,
			expected: "./storage/uploads",
		},
		{
			name:     "Test AppReadTimeout Env",
			value:    appconfigs.Instance.AppReadTimeout,
			expected: 10,
		},
		{
			name:     "Test AppReadTimeout Env",
			value:    appconfigs.Instance.AppUploadLimit,
			expected: 1,
		},
		{
			name:     "TestUpdateEnv Function",
			expected: "UPDATE_SUCCESS",
		},
		{
			name:     "TestUpdateEnv Function ShouldError ReadWrite",
			expected: "UPDATE_ERROR",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			switch test.expected {
			case "UPDATE_SUCCESS":
				initialValue := appconfigs.Instance.AppDebug
				appconfigs.Instance.UpdateEnv("APP_DEBUG", !initialValue)
				assert.Equal(t, appconfigs.Instance.AppDebug, false)
				appconfigs.Instance.UpdateEnv("APP_DEBUG", initialValue)
			case "UPDATE_ERROR":
				viper.Reset()
				initialValue := appconfigs.Instance.AppDebug
				appconfigs.Instance.UpdateEnv("APP_DEBUG", !initialValue)
				assert.Equal(t, appconfigs.Instance.AppDebug, initialValue)
			default:
				assert.Equal(t, test.expected, test.value)
			}
		})
	}
}
