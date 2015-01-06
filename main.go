package main

import (
	"net/http"
	"imgscale"
	"fmt"
	"github.com/go-martini/martini"
)

func main() {
	// Martini
	app := martini.Classic()
	app.Use(imgscale.Middleware("./config/formats.json"))
	http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)
	// http.HandleFunc
	/*
	http.HandleFunc("/", middleware.Middleware("./config/formats.json"))
	http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8080), nil)*/
}

