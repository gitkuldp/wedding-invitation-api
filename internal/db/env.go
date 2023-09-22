package db

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv           string `mapstructure:"APP_ENV"`
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	ServerPort       int    `mapstructure:"SERVER_PORT"`
	ContextTimeout   int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBDialect        string `mapstructure:"DB_DIALECT"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           string `mapstructure:"DB_PORT"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`

	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	MailServer             string `mapstructure:"MAIL_SERVER"`
	MailServerPort         int    `mapstructure:"MAIL_SERVER_PORT"`
	MailUsername           string `mapstructure:"MAIL_USERNAME"`
	MailPass               string `mapstructure:"MAIL_PASS"`
	ClientBaseUrl          string `mapstructure:"CLIENT_BASE_URL_DEV"`
}

func NewEnv() *Env {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	var env Env
	err = viper.Unmarshal(&env)

	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return &env
}
