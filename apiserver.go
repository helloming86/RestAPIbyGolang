package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"miMallDemo/config"
	"miMallDemo/logger"
	"miMallDemo/model"
	"miMallDemo/router"
	v "miMallDemo/version"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg     = pflag.StringP("config", "c", "", "APIServer config file path") // cfg 是一个 *string 指针变量
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	// 使用pflag 进行 flag 绑定（也可以使用官方标准库的flag）
	pflag.Parse()

	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", " ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshalled))
		return
	}

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
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	// 如果提供了 TLS 证书和私钥则启动 HTTPS 端口。
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")

	if cert != "" && key != "" {
		go func() {
			log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("tls.addr"))
			log.Printf(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}

	// 启动HTTP端口
	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
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
