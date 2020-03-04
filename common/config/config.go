package config

import (
	"os"
	"path/filepath"
	"runtime"
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

// Setup initializes the main configuration.
func Setup() {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		panic("config: unable to determine the caller.")
	}
	baseDir := filepath.Dir(b)
	viper.SetConfigName(getFilename())
	viper.AddConfigPath(".")
	viper.AddConfigPath(baseDir)
	viper.AddConfigPath("/etc/itstats")
	viper.AddConfigPath("$HOME/.itstats")
	viper.Set("base_dir", filepath.Dir(baseDir)+"/")
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
