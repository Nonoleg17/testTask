package main

import (
	"testCase/config"
	"testCase/internal/app"
)

func main() {
	cfg := config.NewConfig()
	app.Run(cfg)
}
