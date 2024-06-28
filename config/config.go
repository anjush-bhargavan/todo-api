package config

import "github.com/spf13/viper"

// Config struct maps to the environment variables
type Config struct {
    DBHost     string `mapstructure:"DBHOST"`
    DBKeyspace string `mapstructure:"DBKEYSPACE"`
}

// LoadConfig will load the environment variables to accessible.
func LoadConfig() *Config {
	var cnfg Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&cnfg)
	return &cnfg
}