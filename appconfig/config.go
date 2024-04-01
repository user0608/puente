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
	DataPath() string
	DBLogsLevel() string
}
type config struct {
	ListenAddressValue string `mapstructure:"listen_address"`
	DataPathValue      string `mapstructure:"data_path"`
	DBLogsLevelValue   string `mapstructure:"db_logs_level"`
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
func (c *config) DataPath() string {
	if c.DataPathValue == "" {
		c.DataPathValue = "./data"
	}
	return c.DataPathValue
}
func (c *config) DBLogsLevel() string {
	var valids = map[string]bool{"error": true, "warn": true, "info": true}
	if !valids[c.DBLogsLevelValue] {
		c.DBLogsLevelValue = "silent"
	}
	return c.DBLogsLevelValue
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
