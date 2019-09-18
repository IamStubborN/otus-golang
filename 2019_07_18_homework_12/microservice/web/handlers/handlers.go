package handlers

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Fatal(err)
	}
}