package config

import (
	"github.com/spf13/viper"
)

type config struct {
	RabbitMQUri string `mapstructure:"RABBITMQ_URI"`
}

// Config contains all the env variable
var Config *config

// LoadConfig Load
func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Config)
	return
}
