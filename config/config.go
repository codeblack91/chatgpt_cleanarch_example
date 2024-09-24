package config

import (
	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	}
	Database struct {
		Driver     string `mapstructure:"driver"`
		Connection string `mapstructure:"connection"`
	}
	
}

// LoadConfig loads configuration from a file
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/Users/v/Projects/CleanArchTest/config")

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
