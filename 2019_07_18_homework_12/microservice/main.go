package main

import (
	"fmt"
	"github.com/IamStubborN/otus-golang/2019_07_18_homework_12/microservice/cmd/config"
	"github.com/IamStubborN/otus-golang/2019_07_18_homework_12/microservice/cmd/logger"
	"net/http"

	"github.com/IamStubborN/otus-golang/2019_07_18_homework_12/microservice/web"
)

func main() {
	log := logger.GetLogger()
	cfg := config.GetConfig()
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), web.GetRouter()); err != nil {
		log.Fatal(err)
	}
}
