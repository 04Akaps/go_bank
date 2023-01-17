package util

import (
	"time"

	"github.com/spf13/viper"
)


type Config struct {
	DBDRiver string 	`mapstructure:"DB_DRIVER"`
	DBSource string		`mapstructure:"DB_SOURCE"`
	ServerAddress string`mapstructure:"SERVER_ADDRESS"`
	TokenKey	string	`mapstructure:"TOKEN_KEY"`
	TokenDuration time.Duration `mapstructure:"TOKEN_DURATION"`
}


func LoadConfig(path string) (config  Config ,err error) {
	// path is find config file 
	viper.AddConfigPath(path)
	viper.SetConfigFile("app")
	viper.SetConfigType("env")


	viper.AutomaticEnv()

	err = viper.ReadConfig()

	if err != nil{
		return
	}

	err = viper.Unmarshal(&config)
	return
}