package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Token  string `env:"TOKEN,required" `
	DBPath string `env:"DB_PATH,required"`
}

func New() (*Config, error) {
	c, err := env.ParseAs[Config]()
	if err != nil {
		return nil, err
	}

	return &c, nil
}
