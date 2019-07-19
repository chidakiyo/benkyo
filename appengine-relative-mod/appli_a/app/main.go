package main

import (
	"github.com/chidakiyo/benkyo/appengine-relative-mod/lib_a"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		target := lib_a.TimeString() + " ですよ"
		writer.Write([]byte(target))
	})
	appengine.Main()
}
