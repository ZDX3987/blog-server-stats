package config

import (
	"errors"
	"fmt"
)

func (cfg *Config) Validate() error {
	if cfg.App.Name == "" {
		return errors.New("app.name is required")
	}

	if cfg.App.Port <= 0 || cfg.App.Port > 65535 {
		return fmt.Errorf("invalid app.port: %d", cfg.App.Port)
	}

	if cfg.MySQL.Host == "" {
		return errors.New("mysql.host is required")
	}

	if cfg.MySQL.Port <= 0 || cfg.MySQL.Port > 65535 {
		return fmt.Errorf("invalid mysql.port: %d", cfg.MySQL.Port)
	}

	if cfg.MySQL.Username == "" {
		return errors.New("mysql.username is required")
	}

	if cfg.MySQL.Database == "" {
		return errors.New("mysql.database is required")
	}

	if cfg.Redis.Addr == "" {
		return errors.New("redis.addr is required")
	}

	return nil
}
