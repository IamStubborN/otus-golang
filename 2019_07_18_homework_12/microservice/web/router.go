package web

import (
	"github.com/IamStubborN/otus-golang/2019_07_18_homework_12/microservice/web/handlers"
	"github.com/go-chi/chi"
)

func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", handlers.Hello)
	return router
}
