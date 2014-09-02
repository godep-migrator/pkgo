package main

import (
	"flag"
	"github.com/subosito/pkgo"
	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web/middleware"
	"log"
	"net/http"
)

func init() {
	bind.WithFlag()
}

func main() {
	var err error

	if !flag.Parsed() {
		flag.Parse()
	}

	pkgo.Initialize()

	m := pkgo.NewMux()
	m.Use(middleware.RequestID)
	m.Use(middleware.Logger)
	m.Use(middleware.Recoverer)
	m.Use(middleware.AutomaticOptions)

	http.Handle("/", m)

	listener := bind.Default()
	log.Println("Starting on:", listener.Addr())

	graceful.HandleSignals()
	bind.Ready()

	err = graceful.Serve(listener, http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}

	graceful.Wait()
}
