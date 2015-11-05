package app

import (
	"github.com/spf13/viper"
	"fmt"
	"os"
)

func GetConfig() (*viper.Viper) {

	env := os.Getenv("VIPER_ENV")

	if len(env) == 0 {
		env = "development"
	}

	config := viper.New()

	config.SetConfigName(env) // name of config file (without extension)

	configPath := os.Getenv("VIPER_CONFIG")

	if len(configPath) == 0 {
		configPath = "config"
	}

	config.AddConfigPath(configPath)

	err := config.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
	    panic(fmt.Errorf("Fatal error config file [%s]: %s \n", configPath, err))
	}

	return config
}