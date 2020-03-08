package api

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is the api configuration struct.
type Config struct {
	Host    string
	Port    uint
	IsProd  bool
	LogPath string
}

// GetURL returns the API URL
func (c *Config) GetURL() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// setDefaults sets the default values
func setDefaults() {
	viper.SetDefault("api_host", "localhost")
	viper.SetDefault("api_port", 8080)
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		Host:    viper.GetString("api_host"),
		Port:    viper.GetUint("api_port"),
		IsProd:  viper.GetBool("is_prod"),
		LogPath: viper.GetString("base_dir") + viper.GetString("log_path"),
	}
	return config
}
