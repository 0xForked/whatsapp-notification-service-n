package configs_test

import (
	"github.com/aasumitro/gowa/configs"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppConfig_InitDbConnShouldErrorOpen(t *testing.T) {
	t.Skip()
	viper.Reset()
	viper.SetConfigFile("../.example.env")
	viper.SetConfigType("dotenv")
	configs.LoadEnv()
	configs.DbPool = nil
	configs.Instance.DBDsnURL = "lorem.db"
	configs.Instance.DBDriver = "lorem"
	assert.Panics(t, func() {
		configs.Instance.InitDbConn()
	})
}

func TestAppConfig(t *testing.T) {
	viper.Reset()
	viper.SetConfigFile("../.example.env")
	viper.SetConfigType("dotenv")
	configs.LoadEnv()

	tt := []struct {
		name     string
		value    any
		expected any
		reflect  any
	}{
		{
			name:     "Test AppName Env",
			value:    configs.Instance.AppName,
			expected: "Gowans",
		},
		{
			name:     "Test AppVersion Env",
			value:    configs.Instance.AppVersion,
			expected: "0.0.1-dev",
		},
		{
			name:     "Test AppUrl Env",
			value:    configs.Instance.AppURL,
			expected: "localhost:8000",
		},
		{
			name:     "Test AppUploadPath Env",
			value:    configs.Instance.AppUploadPath,
			expected: "./storage/uploads",
		},
		{
			name:     "Test AppReadTimeout Env",
			value:    configs.Instance.AppReadTimeout,
			expected: 10,
		},
		{
			name:     "Test AppReadTimeout Env",
			value:    configs.Instance.AppUploadLimit,
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
		{
			name:     "TestInitDBConn",
			expected: "DB_CONN",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			switch test.expected {
			case "DB_CONN":
				configs.Instance.DBDsnURL = "../db/local-data.db"
				configs.Instance.DBDriver = "sqlite3"
				configs.Instance.InitDbConn()
				assert.NotEqual(t, configs.DbPool, nil)
			case "UPDATE_SUCCESS":
				initialValue := configs.Instance.AppDebug
				configs.Instance.UpdateEnv("APP_DEBUG", !initialValue)
				assert.Equal(t, configs.Instance.AppDebug, true)
				configs.Instance.UpdateEnv("APP_DEBUG", initialValue)
			case "UPDATE_ERROR":
				viper.Reset()
				initialValue := configs.Instance.AppDebug
				configs.Instance.UpdateEnv("APP_DEBUG", !initialValue)
				assert.Equal(t, configs.Instance.AppDebug, initialValue)
			default:
				assert.Equal(t, test.expected, test.value)
			}
		})
	}
}
