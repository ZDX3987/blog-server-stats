package bootstrap

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"zhangdx.cn/blog-server-stats/internal/config"
	"zhangdx.cn/blog-server-stats/internal/infra"
	"zhangdx.cn/blog-server-stats/internal/job"
	"zhangdx.cn/blog-server-stats/internal/readcount"
	"zhangdx.cn/blog-server-stats/internal/router"
)

func NewApp() error {
	var (
		env        string
		configPath string
	)
	// go run ./cmd/server -env prod -config ./config/config.dev.yaml
	flag.StringVar(&env, "env", "local", "runtime environment")
	flag.StringVar(&configPath, "config", "", "config file path")
	flag.Parse()
	cfg, err := config.Load(configPath, env)
	if err != nil {
		log.Fatalf("load config failed: %v", err)
		return err
	}
	redisClient := NewRedisClient(cfg.Redis)
	mysqlClient := NewMySQLClient(cfg.MySQL)
	redisOperator := infra.NewRedisOperator(redisClient)
	readCountRepository := readcount.NewRepository(mysqlClient)
	readCountService := readcount.NewReadCountService(redisOperator, readCountRepository)
	readCountHandler := readcount.NewReadCountHandler(readCountService)

	syncJob := job.NewReadCountSyncJob(redisOperator, readCountRepository)
	syncJob.Start(context.Background())

	r := gin.Default()
	setGinMode(env)
	router.Init(r, readCountHandler)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func setGinMode(env string) {
	switch env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
