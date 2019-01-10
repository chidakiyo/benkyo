package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/delay"
	"google.golang.org/appengine/log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/delay", delayHandle)
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

var delayA *delay.Function
var delayB *delay.Function

// initで初期化する
func init() {
	delayA = delay.Func("delayA", func(c context.Context, i int64) {
		log.Infof(c, "run - A")
		i = i * 2
		log.Infof(c, "call - B")
		delayB.Call(c, fmt.Sprintf("%d", i))
	})
	delayB = delay.Func("delayB", func(c context.Context, s string) {
		log.Infof(c, "run - B")
		s = s + " points!"
		log.Infof(c, "%s", s)
	})
}

func delayHandle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Infof(c, "call - A")
	delayA.Call(c, int64(10))
	w.Write([]byte("okay!"))
}
