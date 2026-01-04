package config

import (
	"log"
	"time"
	"github.com/spf13/viper"
)

type Config struct {
	APP_ENV string
	Server  ServerConfig
	DB      DBConfig
	JWT     JWTConfig
}

type ServerConfig struct {
	Port string
}

type DBConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	Name         string
	SSLMode      string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration
}

type JWTConfig struct{
	Secret string
	Expiryhour int
}

func Load()* Config{
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()
	viper.SetDefault("SERVER_PORT","8080")
	viper.SetDefault("DB_SSLMODE","disable")
	viper.SetDefault("MaxOpenConns",10)
	viper.SetDefault("MaxIdleConns",5)
	viper.SetDefault("MaxIdleTime","15m")

	if err:= viper.ReadInConfig();err!=nil{
		log.Println("No config file found using env var")
	}

	cfg:= &Config{
		APP_ENV: viper.GetString("APP_ENV"),
		Server: ServerConfig{
			Port: viper.GetString("SERVER_PORT"),
		},
		DB: DBConfig{
			Host : viper.GetString("DB_HOST"),
			Port: viper.GetString("DB_PORT"),
			User: viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name: viper.GetString("DB_NAME"),
			SSLMode: viper.GetString("DB_SSLMODE"),
			MaxOpenConns: viper.GetInt("DB_MAX_OPEN_CONNS"),
			MaxIdleConns: viper.GetInt("DB_MAX_IDLE_CONNS"),
			MaxIdleTime: viper.GetDuration("DB_MAX_IDLE_TIME"),

		},
		JWT: JWTConfig{
			Secret: viper.GetString("JWT_SECRET"),
			Expiryhour: viper.GetInt("JWT_EXPIRTY_HOUR"),
		},
	}
	return cfg
}

