package bootstrap

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"zhangdx.cn/blog-server-stats/internal/config"
)

func NewMySQLClient(c config.MySQLConfig) *sql.DB {
	db, err := sql.Open("mysql", dsn(c))
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxLifetime(c.ConnMaxLifetime)

	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func dsn(c config.MySQLConfig) string {
	params := url.Values{}
	params.Set("charset", c.Charset)
	params.Set("parseTime", "True")
	params.Set("loc", "Local")

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		c.Username, c.Password, c.Host, c.Port, c.Database, params.Encode(),
	)
}
