package tgbot

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Bot struct {
		Token string `yaml:"token"`
		WHURL int    `yaml:"whurl"`
		Port  int    `yaml:"port" default:"80"`
	}
}

func InitConfig(filepath string) (*Config, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = yaml.Unmarshal(file, cfg)
	return cfg, err
}
