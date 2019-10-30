package main

import (
	"go-restful-api/routes"
	_ "go-restful-api/routes/v1"
	_ "go-restful-api/routes/v2"
	"net/http"
	"time"
)

func main() {
	router := routes.Router

	s := &http.Server{
		Addr:           ":8088",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
