package handler

import (
	"fmt"
	"net/http"
	"os"
)

func EnvHandler(w http.ResponseWriter, req *http.Request) {
	u := req.RequestURI
	if name, ok := os.LookupEnv("NAME"); !ok {
		fmt.Fprintf(w, "Name : [%s]\nPath : %s", "not found", u)
	} else {
		fmt.Fprintf(w, "Name : [%s]\nPath : %s", name, u)
	}
}
