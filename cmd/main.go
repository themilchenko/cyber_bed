package main

import (
	"cyber_bed/internal/app"
	"cyber_bed/internal/config"
	"flag"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	var configPath string
	config.ParseFlag(&configPath)
	flag.Parse()

	cfg := config.New()
	if err := cfg.Open(configPath); err != nil {
		log.Fatal("Failed to open config file")
	}

	e := echo.New()
	app := app.New(e, cfg)
	if err := app.Start(); err != nil {
		app.Echo.Logger.Error(err)
	}
}
