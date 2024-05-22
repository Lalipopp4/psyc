package cache

import (
	"os"
	"time"

	"psyc/internal/errors"

	"gopkg.in/yaml.v2"
)

// App configuration
type Config struct {
	Addr     string        `yaml:"addr"`
	Database int           `yaml:"database"`
	Password string        `yaml:"password"`
	User     string        `yaml:"user"`
	TTLToken time.Duration `yaml:"ttl-token"`
	TTLData  time.Duration `yaml:"ttl-data"`
}

// Inits config by parsing toml file
func InitConfig(cfgfile string) (*Config, error) {
	var cfg = &Config{}
	data, err := os.ReadFile(cfgfile)
	if err != nil {
		return nil, errors.ErrorServer{Msg: errors.ErrParseFile}
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, errors.ErrorServer{Msg: errors.ErrDecodeToml}
	}
	return cfg, nil
}
