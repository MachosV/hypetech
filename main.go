package main

import (
	"log"
	"mux"
	_ "mux"
	"net/http"
	"os"
	"os/signal"
	_ "views"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Println("Signal received, terminating")
		os.Exit(1)
	}()
	bindAddr := os.Getenv("BIND_ADDR")
	if bindAddr == "" {
		log.Fatal("BIND_ADDR environment variable not set")
	}
	mux := mux.GetMux()
	if mux == nil {
		log.Fatal("Mux is nil")
	}
	log.Println("Server up and running", bindAddr)
	http.ListenAndServe(bindAddr, mux)
}
