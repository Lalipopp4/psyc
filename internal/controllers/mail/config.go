package mail

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mail struct {
		mail     string
		password string
		port     string
		host     string
	}
}

func InitConfig(filepath string) (*Config, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	return cfg, yaml.Unmarshal(file, cfg)
}
