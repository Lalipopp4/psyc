package transport

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Addr    string        `default:"0.0.0.0:8000" yaml:"addr"`
	Timeout time.Duration `default:"10000" yaml:"timeout"`
}

func InitConfig(filepath string) (*Config, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	return cfg, yaml.Unmarshal(file, cfg)
}
