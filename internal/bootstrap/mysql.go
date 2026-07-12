package bootstrap

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLClient() *sql.DB {
	dsn := "zhangdx:Zz196810486@tcp(rm-wz9y4w067z3i4u1tr1o.mysql.rds.aliyuncs.com)/zhangdx_blog?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
