package main

import (
	"log"
	"net/http"

	chathttp "github.com/beruangcoklat/live-chat/chat/delivery/http"
	"github.com/beruangcoklat/live-chat/config"
	"github.com/gorilla/mux"
)

func runHTTP() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	initUsecase()

	router := mux.NewRouter()
	chathttp.New(router, chatUc)

	port := config.GetConfig().Port
	log.Print("listen :" + port)
	http.ListenAndServe(":"+port, router)
}
