package config

import (
	"api-repository/pkg/db/minio"
	"api-repository/pkg/db/postgres"
	"api-repository/pkg/db/redis"

	"github.com/ilyakaznacheev/cleanenv"
)

type MainConfig struct {
	GatewayPort     int               `yaml:"GATEWAY_PORT" env:"GATEWAY_PORT" env-default:"8080"`
	UserServicePort int               `yaml:"USER_SERVICE_PORT" env:"USER_SERVICE_PORT" env-default:"50051"`
	UserServiceAddr string            `yaml:"USER_SERVICE_ADDR" env:"SOURCE_SERVICE_ADDR" env-default:"localhost:50051"`
	FileServicePort int               `yaml:"FILE_SERVICE_PORT" env:"FILE_SERVICE_PORT" env-default:"50052"`
	POSTGRES        postgres.PgConfig `yaml:"POSTGRES"`
	REDIS           redis.RConfig     `yaml:"REDIS"`
	MinIO           minio.MnConfig    `yaml:"MINIO"`
}

func NewMainConfig() (*MainConfig, error) {
	var cfg MainConfig

	if err := cleanenv.ReadConfig("config/config.yml", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
