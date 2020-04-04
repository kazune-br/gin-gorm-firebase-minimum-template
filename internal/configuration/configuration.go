package configuration

import (
	"github.com/kelseyhightower/envconfig"
)

// Config is environment variables list
type Config struct {
	Env            string `envconfig:"ENV"`
	AppPort        int    `envconfig:"APP_PORT"`
	DBUser         string `envconfig:"DB_USER"`
	DBPass         string `envconfig:"DB_PASS"`
	DBHost         string `envconfig:"DB_HOST"`
	DBPort         string `envconfig:"DB_PORT"`
	DBName         string `envconfig:"DB_NAME"`
	FireBaseJson64 string `envconfig:"FIREBASE_JSON64"`
}

var globalConfig Config

func Load() {
	envconfig.MustProcess("", &globalConfig)
}

func Get() Config {
	return globalConfig
}
