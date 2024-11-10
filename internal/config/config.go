package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var dbEnvs = []string{"ONFLY_DB_USER", "ONFLY_DB_PASS", "ONFLY_DB_HOST", "ONFLY_DB_PORT", "ONFLY_DB_NAME", "TIME_ZONE"}
var env = fmt.Sprintf(".env.%s", os.Getenv("ENV"))

type Config struct {
	DbUser   string `mapstructure:"ONFLY_DB_USER"`
	DbPass   string `mapstructure:"ONFLY_DB_PASS"`
	DbHost   string `mapstructure:"ONFLY_DB_HOST"`
	DbPort   string `mapstructure:"ONFLY_DB_PORT"`
	DbName   string `mapstructure:"ONFLY_DB_NAME"`
	TimeZone string `mapstructure:"TIME_ZONE"`
}

func InitConfig() (Config, error) {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	var config Config

	if err := godotenv.Load(env); err != nil {
		return config, err
	}

	env := fmt.Sprintf(".env.%s", os.Getenv("ENV"))

	viper.AddConfigPath("./internal/config")
	viper.SetConfigName(env)

	for _, env := range dbEnvs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
