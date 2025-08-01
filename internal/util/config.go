package util

import "github.com/spf13/viper"

type Config struct {
	APPPROT string `mapstructure:"APP_PRORT"`
	APPNAME string `mapstructure:"APP_NAME"`
	APPDEBUG string `mapstructure:"APP_DEBUG"`
}

func LoadConfig(path string) (config Config , err error) {

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if(err != nil) {
		return
	}

	err = viper.Unmarshal(&config)
	return
}