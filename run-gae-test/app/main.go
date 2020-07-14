package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	p := os.Getenv("PORT")
	if len(p) == 0 {
		p = "8080"
	}

	// 「/a」に対して処理を追加
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		host := r.Host
		fmt.Fprintf(w, "(%s) (%s)", host, path)
	})

	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)

}