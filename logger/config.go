package logger

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is the logger configuration struct.
type Config struct {
	FilePath string
	IsDebug  bool
}

// NewConfig returns the configuration object.
func NewConfig() (*Config, error) {
	config := &Config{
		FilePath: viper.GetString("config_file_path"),
		IsDebug:  viper.GetBool("config_is_debug"),
	}
	if config.FilePath == "" {
		return nil, fmt.Errorf("config_file_path must be set")
	}
	return config, nil
}
