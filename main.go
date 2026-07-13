package main

import (
	"log"

	"zhangdx.cn/blog-server-stats/internal/bootstrap"
)

func main() {
	if err := bootstrap.NewApp(); err != nil {
		log.Fatalf("app bootstrap failed: %v", err)
	}
}
