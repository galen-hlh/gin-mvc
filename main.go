package main

import (
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"go-restful-api/app/compoments/config"
	"go-restful-api/app/compoments/mysql"
	"go-restful-api/app/compoments/redis"
	"go-restful-api/routes"
	_ "go-restful-api/routes/v1"
	_ "go-restful-api/routes/v2"
	"net/http"
	"time"
)

func main() {

	// 初始化配置
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("init config failure")
		return
	}
	err = redis.InitRedis(cfg.RedisConfig)
	if err != nil {
		fmt.Println("init redis failure")
		return
	}
	err = mysql.InitDB(cfg.DBConfig)
	if err != nil {
		fmt.Println("init mysql failure")
		return
	}

	router := routes.Router
	ginpprof.Wrapper(router)
	s := &http.Server{
		Addr:           ":8088",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
