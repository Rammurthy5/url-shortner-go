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
		HTTPConfig  HTTPConfig  `mapstructure:"http"`
		DBConfig    DBConfig    `mapstructure:"db"`
		CacheConfig CacheConfig `mapstructure:"cache"`
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
	CacheConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"uname"`
		Password string `mapstructure:"password"`
		Db       int    `mapstructure:"db"`
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
	// db config bind
	envConfig.BindEnv("db.host", "db_host")
	envConfig.BindEnv("db.port", "db_port")
	envConfig.BindEnv("db.username", "db_username")
	envConfig.BindEnv("db.password", "db_password")
	envConfig.BindEnv("db.dbname", "dbname")

	// cache config bind
	envConfig.BindEnv("cache.host", "cache_host")
	envConfig.BindEnv("cache.port", "cache_port")
	envConfig.BindEnv("cache.username", "cache_uname")
	envConfig.BindEnv("cache.password", "cache_password")
	envConfig.BindEnv("cache.db", "cache_db")

	if err := envConfig.Unmarshal(&c); err != nil {
		log.Fatal(fmt.Errorf("fatal error unmarshaling config file: %w ", err))
	}
	return c, nil
}
