package main

import "miMallDemo/router"

func main() {
	r := router.SetupRouter()
	r.Run()
}
