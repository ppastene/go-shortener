package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host            string `mapstructure:"HOST"`
	Port            int    `mapstructure:"PORT"`
	CacheService    string `mapstructure:"CACHE_SERVICE"`
	ExpirationTime  int    `mapstructure:"EXPIRATION_TIME"`
	KeyLength       int    `mapstructure:"KEY_LENGTH"`
	RedisHost       string `mapstructure:"REDIS_HOST"`
	RedisPort       int    `mapstructure:"REDIS_PORT"`
	RedisPassword   string `mapstructure:"REDIS_PASSWORD"`
	RedisDB         string `mapstructure:"REDIS_DB"`
	InternalCleanup int    `mapstructure:"INTERNAL_CLEANUP"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("CACHE_SERVICE", "internal")
	viper.SetDefault("EXPIRATION_TIME", 3600)
	viper.SetDefault("KEY_LENGTH", 6)
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", 6739)
	viper.SetDefault("REDIS_PASSWORD", "password")
	viper.SetDefault("REDIS_DB", "0")
	viper.SetDefault("INTERNAL_CLEANUP", 0)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found. Setting the config variables to default")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
