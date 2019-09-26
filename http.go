package main

import (
	"net/http"
	"t-gin/routes"
	_ "t-gin/routes/v1"
	_ "t-gin/routes/v2"
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
