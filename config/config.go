package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

const prefix = "its"

// getFilename returns filename based on the environment variable
func getFilename() string {
	fileName := "config"
	env := os.Getenv(strings.ToUpper(prefix + "_env"))
	if env != "" {
		fileName += "." + env
	}
	return fileName
}

// Init initializes the main configuration.
func Setup() {
	viper.SetConfigName(getFilename())
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/itstats")
	viper.AddConfigPath("$HOME/.itstats")
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
