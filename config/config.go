package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type (
	Config struct {
		HTTPConfig HTTPConfig `mapstructure:"http"`
		DBConfig   DBConfig   `mapstructure:"db"`
	}
	HTTPConfig struct {
		Port string `mapstructure:"port"`
	}
	DBConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
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
		log.Fatal(fmt.Errorf("fatal error config file: %w ", err))
	}

	envConfig := viper.Sub(viper.GetString("env"))
	if err := envConfig.Unmarshal(&c); err != nil {
		log.Fatal(fmt.Errorf("fatal error unmarshaling config file: %w ", err))
	}
	return c, nil
}
