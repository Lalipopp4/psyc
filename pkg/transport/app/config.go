package transport

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Addr    string `default:"0.0.0.0:8000" yaml:"addr"`
		Timeout int    `default:"10" yaml:"timeout"`
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
