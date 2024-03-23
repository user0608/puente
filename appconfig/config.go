package appconfig

import (
	"errors"
	"log/slog"
	"path"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig interface {
	StoreDir() string
}
type config struct {
	StoreDirValue string `mapstructure:"store_dir"`
}

func (c *config) validate() error {
	if c.StoreDirValue == "" {
		slog.Error("store_dir not found con config file")
		return errors.New("store_dir nof found")
	}
	return nil
}

func (c *config) StoreDir() string {
	return c.StoreDirValue
}

func LoadAppConfig(configPath string) (AppConfig, error) {
	configPath = strings.TrimSpace(configPath)
	var c config
	if configPath != "" {
		v := viper.New()
		v.SetConfigFile(configPath)
		if err := v.Unmarshal(&c); err != nil {
			slog.Error("unmarshal config", "file", path.Base(configPath), "error", err)
			return nil, err
		}
		if err := c.validate(); err != nil {
			return nil, err
		}
		return &c, nil
	}
	// todo: try load from envs
	if err := c.validate(); err != nil {
		return nil, err
	}
	return &c, nil
}
