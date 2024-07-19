package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type ApiConfig struct {
	Port       int    `env:"PORT, default=5000"`
	EnvName    string `env:"ENV_NAME, default=development"`
	ApiVersion string `env:"VERSION, default=v1"`
}

func (c *ApiConfig) IsDevelopment() bool {
	return c.EnvName == "development"
}

type DatabaseConfig struct {
	Url           string `env:"URL, required"`
	UrlSecretName string `env:"URL_SECRET_NAME, required"`
}

type CloudConfig struct {
	BaseEndpoint string `env:"BASE_ENDPOINT"`
}

func (c *CloudConfig) IsBaseEndpointSet() bool {
	return c.BaseEndpoint != ""
}

type CacheConfig struct {
	Host           string `env:"HOST, required"`
	HostSecretName string `env:"HOST_SECRET_NAME, required"`
	Password       string `env:"PASSWORD"`
	DB             int    `env:"DB, default=0"`
}

type Config struct {
	ApiConfig   *ApiConfig      `env:",prefix=API_"`
	DbConfig    *DatabaseConfig `env:",prefix=DB_"`
	CloudConfig *CloudConfig    `env:",prefix=AWS_"`
	CacheConfig *CacheConfig    `env:",prefix=CACHE_"`
}

func LoadFromEnv(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
