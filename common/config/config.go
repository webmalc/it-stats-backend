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

// setDefaults sets default values
func setDefaults(baseDir string) {
	viper.Set("base_dir", filepath.Dir(baseDir)+"/")
	viper.SetDefault("is_prod", false)
	viper.SetDefault("log_path", "logs/app.log")
}

// setPaths set paths
func setPaths(baseDir string) {
	viper.AddConfigPath(".")
	viper.AddConfigPath(baseDir)
	viper.AddConfigPath("/etc/itstats")
	viper.AddConfigPath("$HOME/.itstats")
}

// setEnv sets the environment
func setEnv() {
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
}

// Setup initializes the main configuration.
func Setup() {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		panic("config: unable to determine the caller.")
	}
	baseDir := filepath.Dir(b)
	viper.SetConfigName(getFilename())
	setPaths(baseDir)
	setEnv()
	setDefaults(baseDir)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
