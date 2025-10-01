package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var serviceConfig ServiceConfig

func init() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("missing configuration path environment variable")
	}
	err := cleanenv.ReadConfig(configPath, &serviceConfig)
	if err != nil {
		panic("configuration read phase failed: " + err.Error())
	}
}

type (
	GrpcConfig struct {
		Port             uint16 `yaml:"port"`
		CertificatesPath string `yaml:"certificates-path"`
	}

	TokenStorageConfig struct {
		Address  string `yaml:"address"`
		Port     uint16 `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
		Protocol int    `yaml:"protocol"`
	}

	UserStorageConfig struct {
		DB       string `yaml:"db"`
		Host     string `yaml:"host"`
		Port     uint16 `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}

	ServiceConfig struct {
		Env          string             `yaml:"env"`
		GRPC         GrpcConfig         `yaml:"grpc"`
		TokenStorage TokenStorageConfig `yaml:"token-storage"`
		UserStorage  UserStorageConfig  `yaml:"user-storage"`
	}
)

func Load() *ServiceConfig {
	return &serviceConfig
}
