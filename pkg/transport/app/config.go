package transport

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Addr    string `yaml:"addr"`
	Timeout int    `yaml:"timeout"`
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
