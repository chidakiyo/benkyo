package main

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/_ah/warmup", warmup)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Infof(c, "Hello, world!")
}

func warmup(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Infof(c, "Warm Up!")
}
