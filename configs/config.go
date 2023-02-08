package configs

import (
	"database/sql"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"sync"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type AppConfig struct {
	AppDebug       bool   `mapstructure:"APP_DEBUG"`
	AppName        string `mapstructure:"APP_NAME"`
	AppURL         string `mapstructure:"APP_URL"`
	AppVersion     string `mapstructure:"APP_VERSION"`
	AppUploadPath  string `mapstructure:"APP_UPLOAD_PATH"`
	AppReadTimeout int    `mapstructure:"APP_READ_TIMEOUT"`
	AppUploadLimit int    `mapstructure:"APP_UPLOAD_LIMIT"`
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBDsnURL       string `mapstructure:"DB_DSN_URL"`
	Gpt3APIKey     string `mapstructure:"GPT_API_KEY"`
	BotReqMSISDN   string `mapstructure:"BOT_REQ_MSISDN"`
}

var (
	cfgOnce  sync.Once
	dbOnce   sync.Once
	Instance *AppConfig
	DbPool   *sql.DB
)

func init() {
	// set config file
	viper.SetConfigFile(".env")
}

func LoadEnv() {
	log.Printf("Load configuration file . . . .")
	// find environment file
	viper.AutomaticEnv()
	// read env handler
	readEnv := func() {
		// error handling for specific case
		if err := viper.ReadInConfig(); err != nil {
			// specified error message
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
				log.Println(".env file not found!, please copy .example.env and paste as .env")
			}
			// general error message
			log.Printf("ENV_ERROR: %s", err.Error())
		}
		// extract config to struct
		if err := viper.Unmarshal(&Instance); err != nil {
			panic(fmt.Sprintf("ENV_ERROR: %s", err.Error()))
		}
	}
	// instance
	cfgOnce.Do(func() {
		// read env
		readEnv()
		// subs to event
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Printf("update configuration data . . . .")

			readEnv()
		})
		// watch file update
		viper.WatchConfig()
		// notify that config file is ready
		log.Println("configuration file: ready")
	})
}

func (cfg *AppConfig) UpdateEnv(key, value any) {
	if err := viper.ReadInConfig(); err != nil {
		log.Println("READ", err.Error())
	}

	viper.Set(key.(string), value)

	viper.SetConfigType("dotenv")

	if err := viper.WriteConfig(); err != nil {
		log.Println("WRITE", err.Error())
	}
}

func (cfg *AppConfig) InitDbConn() {
	dbOnce.Do(func() {
		db, err := sql.Open(cfg.DBDriver, cfg.DBDsnURL)
		if err != nil {
			panic(fmt.Sprintf("DATABASE_ERROR: %s", err.Error()))
		}

		DbPool = db

		if err := DbPool.Ping(); err != nil {
			panic(fmt.Sprintf("DATABASE_ERROR: %s", err.Error()))
		}

		log.Printf("Database connection pool created with %s driver . . . .", cfg.DBDriver)
	})
}
