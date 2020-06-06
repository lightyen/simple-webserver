package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"app/server"
)

var port string
var public string

func init() {
	flag.StringVar(&port, "port", "http", "web service port number")
	flag.StringVar(&public, "public", "build", "static web resouces place")
	flag.Parse()
}

func main() {
	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        server.NewRouter(public),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
