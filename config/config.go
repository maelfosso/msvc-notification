package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type mailConfig struct {
	Host     string `env:"MAIL_HOST"`
	Username string `env:"MAIL_USERNAME"`
	Password string `env:"MAIL_PASSWORD"`
	// APIKey   string `env:"MAIL_API_KEY"`
}

// Config is a application configuration structure
type Config struct {
	EnvMode     string `env:"ENV_MODE" env-required:"true"`
	WebAppURL   string `env:"WEBAPP_URL" env-required:"true"`
	RabbitMQUri string `env:"RABBITMQ_URI" env-required:"true"`

	Mail struct {
		Host     string `env:"MAIL_HOST"`
		Port     string `env:"MAIL_PORT"`
		Username string `env:"MAIL_USERNAME"`
		Password string `env:"MAIL_PASSWORD"`
		APIKey   string `env:"MAIL_API_KEY"`
	}
}

// Config contains all the env variable
var config Config

// LoadConfig Load
func LoadConfig(path string) (err error) {
	err = cleanenv.ReadEnv(&config)
	// err = cleanenv.ReadConfig(path, &config)

	if err != nil {
		return
	}

	return
}

// GetConfig returns the App configuration
func GetConfig() Config {
	return config
}
