package appconfig

import (
	"log/slog"
	"path"
	"puente/cmd"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig interface {
	ListenAddress() string
}
type config struct {
	ListenAddressValue string `mapstructure:"listen_address"`
}

func (c *config) validate() error {

	return nil
}

func (c *config) ListenAddress() string {
	if c.ListenAddressValue == "" {
		c.ListenAddressValue = "localhost:10265"
	}
	return c.ListenAddressValue
}

func LoadAppConfig(confPath cmd.ConfigPath) (AppConfig, error) {

	var configPath = strings.TrimSpace(string(confPath))
	var c config
	if configPath != "" {
		v := viper.New()
		v.SetConfigFile(configPath)
		v.SetConfigType("yaml")
		if err := v.ReadInConfig(); err != nil {
			slog.Error("reading config", "file", path.Base(configPath), "error", err)
			return nil, err
		}
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
