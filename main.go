package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"miMallDemo/router"
	"net/http"
	"time"
)

func main() {
	// create the gin engine
	g := gin.New()

	// gin middlewares
	//var middlewares []gin.HandlerFunc
	middlewares := []gin.HandlerFunc{}

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
	log.Printf("Start to listening the incoming requests on http address: %s", ":8080") // 标注2
	log.Printf(http.ListenAndServe(":8080", g).Error())  // 标注3
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	// Ping the server by sending a GET request to `/health`.
	for i := 0; i < 2; i++ {
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep for a second to continue the next ping
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}