package main

import (
	"api-perpus/app"
	"api-perpus/config"
)

func main() {
	config := config.InitConfig()
	app.Run(config)
}
