package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"miMallDemo/config"
	"miMallDemo/logger"
	"miMallDemo/model"
	"miMallDemo/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config","c","","APIServer config file path") // cfg 是一个 *string 指针变量
)

func main() {
	// 使用pflag 进行 flag 绑定（也可以使用官方标准库的flag）
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init Logger
	logger.InitLogger()

	// init DB
	model.InitDB()
	defer model.DB.Close()

	if model.DB.HasTable(&model.User{}) {
		model.DB.AutoMigrate(&model.User{})
	} else {
		model.DB.CreateTable(&model.User{})
	}

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// create the gin engine
	g := gin.New()

	// gin middlewares
	//var middlewares []gin.HandlerFunc
	var middlewares []gin.HandlerFunc

	// routers
	router.Load(
		// cores
		g,
		// middlewares
		middlewares...,
		)

	// Ping the server to make sure the router is working.
	// 注意：go func() 声明在log.Printf()之前 标注1 声明在标注2、3前 为什么？
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}() // 标注1
	// 让main goroutine 暂停运行500ms
	// 这时，标注2和3不会执行，但是 go gunc 会并发继续执行，因为WEB服务还未启动，提示 Waiting for the router, retry in 1 second.
	// main goroutine 暂停运行时间一过，标注2、3继续执行
	// 标注2 执行后，提示：Start to listening the incoming requests on http address: :8080
	// 标注3，成功启动WEB服务
	// 并发 go gunc 在WEB服务启动并没有Error的情况下，提示： The router has been deployed successfully.
	time.Sleep(time.Millisecond * 500)
	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr")) // 标注2
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())  // 标注3
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	// Ping the server by sending a GET request to `/health`.
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep for a second to continue the next ping
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}