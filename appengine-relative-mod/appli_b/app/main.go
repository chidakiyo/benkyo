package main

import (
	"github.com/chidakiyo/benkyo/appengine-relative-mod/lib_b"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		target := lib_b.LowerString("HELLO WORLD!")
		writer.Write([]byte(target))
	})
	appengine.Main()
}
