package util

import "github.com/spf13/viper"

type Config struct {
	AppPort      string `mapstructure:"APP_PORT"`
	AppName      string `mapstructure:"APP_NAME"`
	AppDebug     string `mapstructure:"APP_DEBUG"`

	DBConnection string `mapstructure:"DB_CONNECTION"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBUsername   string `mapstructure:"DB_USERNAME"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBDatabase   string `mapstructure:"DB_DATABASE"`
	DBPort       string `mapstructure:"DB_PORT"`
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