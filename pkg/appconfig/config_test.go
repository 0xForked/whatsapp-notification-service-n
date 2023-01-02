package appconfig_test

import (
	"github.com/aasumitro/gowa/pkg/appconfig"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppConfig(t *testing.T) {
	viper.Reset()
	viper.SetConfigFile("../../.example.env")
	viper.SetConfigType("dotenv")
	appconfig.LoadEnv()

	tt := []struct {
		name     string
		value    any
		expected any
		reflect  any
	}{
		{
			name:     "Test AppName Env",
			value:    appconfig.Instance.AppName,
			expected: "Gowans",
		},
		{
			name:     "Test AppSessionPath Env",
			value:    appconfig.Instance.AppDescription,
			expected: "Whatsapp Notification Service",
		},
		{
			name:     "Test AppVersion Env",
			value:    appconfig.Instance.AppVersion,
			expected: "0.0.1-dev",
		},
		{
			name:     "Test AppUrl Env",
			value:    appconfig.Instance.AppURL,
			expected: "localhost:8000",
		},
		{
			name:     "Test AppUploadPath Env",
			value:    appconfig.Instance.AppUploadPath,
			expected: "./storage/uploads",
		},
		{
			name:     "Test AppReadTimeout Env",
			value:    appconfig.Instance.AppReadTimeout,
			expected: 10,
		},
		{
			name:     "Test AppReadTimeout Env",
			value:    appconfig.Instance.AppUploadLimit,
			expected: 1024,
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
				initialValue := appconfig.Instance.AppDebug
				appconfig.Instance.UpdateEnv("APP_DEBUG", !initialValue)
				assert.Equal(t, appconfig.Instance.AppDebug, true)
				appconfig.Instance.UpdateEnv("APP_DEBUG", initialValue)
			case "UPDATE_ERROR":
				viper.Reset()
				initialValue := appconfig.Instance.AppDebug
				appconfig.Instance.UpdateEnv("APP_DEBUG", !initialValue)
				assert.Equal(t, appconfig.Instance.AppDebug, initialValue)
			default:
				assert.Equal(t, test.expected, test.value)
			}
		})
	}
}
