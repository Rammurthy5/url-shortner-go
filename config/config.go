package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type (
	Config struct {
		HttpConfig HttpConfig `mapstructure:"http"`
	}
	HttpConfig struct {
		Port string `mapstructure:"port"`
	}
)

func Load() (Config, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal(fmt.Println("Error loading .env file:", err))
	}
	var c Config
	//viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	//viper.AddConfigPath("../config/")
	//viper.AddConfigPath("../../config/")
	viper.SetConfigFile("config/config.yaml")
	viper.AutomaticEnv() // Automatically override values with environment variables
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	envConfig := viper.Sub(viper.GetString("env"))
	if err := envConfig.Unmarshal(&c); err != nil {
		log.Fatal(fmt.Errorf("Fatal error unmarshaling config file: %w \n", err))
	}
	return c, nil
}
