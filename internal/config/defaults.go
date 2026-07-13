package config

import "github.com/spf13/viper"

func setDefaults(v *viper.Viper) {
	v.SetDefault("app.name", "blog-api")
	v.SetDefault("app.env", "local")
	v.SetDefault("app.port", 8080)
	v.SetDefault("app.debug", false)

	v.SetDefault("server.read_timeout", "10s")
	v.SetDefault("server.write_timeout", "10s")
	v.SetDefault("server.idle_timeout", "60s")
	v.SetDefault("server.shutdown_timeout", "10s")

	v.SetDefault("log.level", "info")
	v.SetDefault("log.format", "json")
	v.SetDefault("log.output", "stdout")

	v.SetDefault("mysql.charset", "utf8mb4")
	v.SetDefault("mysql.max_open_conns", 50)
	v.SetDefault("mysql.max_idle_conns", 10)
	v.SetDefault("mysql.conn_max_lifetime", "1h")

	v.SetDefault("redis.db", 0)
	v.SetDefault("redis.pool_size", 20)
	v.SetDefault("redis.dial_timeout", "3s")
	v.SetDefault("redis.read_timeout", "2s")
	v.SetDefault("redis.write_timeout", "2s")
}
