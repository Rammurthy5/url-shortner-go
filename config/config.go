package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
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
	if err := envConfig.BindEnv("db.host", "db_host"); err != nil {
		// Handle the error, e.g., log it or return an error
		log.Printf("Failed to bind env variable db.host: %v", err)
	}

	if err := envConfig.BindEnv("db.port", "db_port"); err != nil {
		log.Printf("Failed to bind env variable db.port: %v", err)
	}

	if err := envConfig.BindEnv("db.username", "db_username"); err != nil {
		log.Printf("Failed to bind env variable db.username: %v", err)
	}

	if err := envConfig.BindEnv("db.password", "db_password"); err != nil {
		log.Printf("Failed to bind env variable db.password: %v", err)
	}

	if err := envConfig.BindEnv("db.dbname", "db_dbname"); err != nil {
		log.Printf("Failed to bind env variable db.dbname: %v", err)
	}

	if err := envConfig.BindEnv("cache.host", "cache_host"); err != nil {
		log.Printf("Failed to bind env variable cache.host: %v", err)
	}

	// cache config bind
	if err := envConfig.BindEnv("cache.port", "cache_port"); err != nil {
		log.Printf("Failed to bind env variable cache.port: %v", err)
	}
	if err := envConfig.BindEnv("cache.username", "cache_uname"); err != nil {
		log.Printf("Failed to bind env variable cache.username: %v", err)
	}
	if err := envConfig.BindEnv("cache.password", "cache_password"); err != nil {
		log.Printf("Failed to bind env variable cache.password: %v", err)
	}
	if err := envConfig.BindEnv("cache.db", "cache_db"); err != nil {
		log.Printf("Failed to bind env variable cache.db: %v", err)
	}

	if err := envConfig.Unmarshal(&c); err != nil {
		log.Fatal(fmt.Errorf("fatal error unmarshaling config file: %w ", err))
	}
	return c, nil
}
