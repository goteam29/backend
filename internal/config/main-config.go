package config

import (
	"api-repository/pkg/db/postgres"
	"api-repository/pkg/db/redis"
	"github.com/ilyakaznacheev/cleanenv"
)

type MainConfig struct {
	GatewayPort     int `yaml:"GATEWAY_PORT" env:"GATEWAY_PORT" env-default:"8080"`
	UserServicePort int `yaml:"USER_SERVICE_PORT" env:"USER_SERVICE_PORT" env-default:"50050"`
	PgConf          postgres.PgConfig
	RedisConf       redis.RConfig
}

func NewMainConfig() (*MainConfig, error) {
	var cfg MainConfig

	if err := cleanenv.ReadConfig("./config/config.yml", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
