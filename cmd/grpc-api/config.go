package main

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Address      string        `default:":8080" required:"true"`
		ReadTimeout  time.Duration `default:"10s" required:"true"`
		WriteTimeout time.Duration `default:"10s" required:"true"`
	}

	Redis struct {
		Host     string `default:"127.0.0.1"`
		Port     int16  `default:"6379"`
		Password string `default:"notasecret"`
	}
}

func ParseConfig() (*Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return nil, fmt.Errorf("config: failed to parse from env, %v", err)
	}

	return &config, nil
}

func (config *Config) RedisAddr() string {
	return fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
}
