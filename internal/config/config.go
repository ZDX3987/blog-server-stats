package config

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App    AppConfig    `mapstructure:"app"`
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Log    LogConfig    `mapstructure:"log"`
}

type AppConfig struct {
	Name  string `mapstructure:"name"`
	Env   string `mapstructure:"env"`
	Port  int    `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

type ServerConfig struct {
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Output string `mapstructure:"output"`
}

type MySQLConfig struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	Database        string        `mapstructure:"database"`
	Charset         string        `mapstructure:"charset"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

type RedisConfig struct {
	Addr         string        `mapstructure:"addr"`
	Password     string        `mapstructure:"password"`
	DB           int           `mapstructure:"db"`
	PoolSize     int           `mapstructure:"pool_size"`
	DialTimeout  time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

func Load(configPath, env string) (*Config, error) {
	log.Printf("Load config from path: %s, env: %s", configPath, env)
	v := viper.New()
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		v.SetConfigName("config." + env)
		v.SetConfigType("yml")
		v.AddConfigPath("./config")
		v.AddConfigPath(".")
	}
	//setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}
	// 允许环境变量覆盖配置文件
	v.AutomaticEnv()
	// 设置环境变量键转换 mysql.host => MYSQL_HOST
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}
	return &cfg, nil
}
