package config

import (
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/L1LSunflower/auction/pkg/logger/message"
)

type Config struct {
	AppUrl      string
	AppPort     int
	BearerToken string
	DB          *DB
	Redis       *Redis
	Log         *LogConfig
}

type DB struct {
	DBDriver string
	DBString string
}

type Redis struct {
	Address  string
	Port     string
	Password string
}

type LogConfig struct {
	Level  string
	Driver string
}

var (
	configInstance     *Config
	onceConfigInstance sync.Once
)

func GetConfig() *Config {
	if configInstance == nil {
		onceConfigInstance.Do(func() {

			appUrl := getValue("APP_URL", "")
			appUrl = strings.TrimRight(appUrl, "/")

			appPort := 3000
			if p, err := strconv.Atoi(getValue("LISTEN_PORT", "3000")); err == nil && p != 0 {
				appPort = p
			}

			configInstance = &Config{
				AppUrl:      appUrl,
				AppPort:     appPort,
				BearerToken: getValue("BEARER_TOKEN", "fb250092-974c-44a7-b4ed-4e71b5875886"),
				DB: &DB{
					DBDriver: getValue("DB_DRIVER", "mysql"),
					DBString: getValue("DB_STRING", "root:secret@tcp(mariadb:3306)/auction?parseTime=true"),
				},
				Redis: &Redis{
					Address:  getValue("REDIS_ADDRESS", "redis"),
					Port:     getValue("REDIS_PORT", "6379"),
					Password: getValue("REDIS_PASSWORD", ""),
				},
				Log: &LogConfig{
					Level:  getValue("LOG_LEVEL", message.InfoLevel),
					Driver: getValue("LOG_DRIVER", logger.DriverStdout),
				},
			}
		})
	}
	return configInstance
}

func getValue(name, defaultValue string) string {
	value := os.Getenv(name)
	if len(value) <= 0 {
		return defaultValue
	}
	return value
}
