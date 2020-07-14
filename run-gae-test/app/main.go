package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	p := os.Getenv("PORT")
	if len(p) == 0 {
		p = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		host := r.Host
		fmt.Fprintf(w, "(%s) (%s)", host, path)
	})

	http.HandleFunc("/exec", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if qq,ok := q["q"]; !ok || len(qq) > 1 {
			fmt.Fprintf(w, "bye")
			return
		} else {

			qqq := strings.Split(qq[0], " ")

			b, err := exec.Command(qqq[0], qqq[1:]...).Output()
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			}
			fmt.Fprintf(w, "%s", b)
		}
	})

	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)

}