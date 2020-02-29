package config

import (
	"github.com/spf13/viper"
)

// Init initializes the main configuration.
func Setup() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/itstats")
	viper.AddConfigPath("$HOME/.itstats")
	viper.SetEnvPrefix("its")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
