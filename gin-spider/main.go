package main

import (
	"gin-spider/router"
)

func main() {
	app := router.App()
	app.Run(":9090")
}
