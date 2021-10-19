package main

import (
	"flag"

	chatusecase "github.com/beruangcoklat/live-chat/chat/usecase"
	"github.com/beruangcoklat/live-chat/config"
	"github.com/beruangcoklat/live-chat/domain"
)

var (
	chatUc domain.ChatUsecase
)

func initConfig() error {
	return config.Init("/etc/live-chat/config.json")
}

func initUsecase() {
	chatUc = chatusecase.New()
}

func main() {
	var app string
	flag.StringVar(&app, "app", "http", "app type (http)")
	flag.Parse()

	switch app {
	case "http":
		runHTTP()
	}
}
