package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"strings"
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
		Database string `mapstructure:"dbname"`
	}
)

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Println("Error loading .env file:", err))
	}
	var c Config
	viper.AddConfigPath("./config")
	viper.SetConfigFile("config/config.yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // Automatically override values with environment variables
	env := viper.GetString("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %w ", err))
	}

	envConfig := viper.Sub(env)
	envConfig.BindEnv("dev.db.host", "host")
	envConfig.BindEnv("dev.db.port", "port")
	envConfig.BindEnv("dev.db.username", "username")
	envConfig.BindEnv("dev.db.password", "password")
	envConfig.BindEnv("dev.db.dbname", "dbname")

	if err := envConfig.Unmarshal(&c); err != nil {
		log.Fatal(fmt.Errorf("fatal error unmarshaling config file: %w ", err))
	}
	return c, nil
}
