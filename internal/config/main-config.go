package config

import (
	fileConfig "api-repository/internal/file-service/service/config"
	userConfig "api-repository/internal/user-sevice/service/config"
	"github.com/ilyakaznacheev/cleanenv"
)

type MainConfig struct {
	FileServiceConfig *fileConfig.Config
	UserServiceConfig *userConfig.Config
}

func NewMainConfig() (*MainConfig, error) {
	var cfg MainConfig

	if err := cleanenv.ReadConfig("./config/config.yml", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
