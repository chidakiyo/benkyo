package main

import (
	"fmt"
	"github.com/chidakiyo/go-sandbox/appengine/datastore"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/delay"
	"google.golang.org/appengine/log"
	"net/http"
)

func main() {
	http.HandleFunc("/delay", delayHandle)
	http.HandleFunc("/_ah/warmup", warmup)

	http.HandleFunc("/ds/write", datastore.DsWrite)
	http.HandleFunc("/ds/read", datastore.DsRead)
	http.HandleFunc("/", handle)

	//// [START setting_port]
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "8080"
	//	l.Printf("Defaulting to port %s", port)
	//}
	//
	//l.Printf("Listening on port %s", port)
	//l.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	//// [END setting_port]

	// listenより前に呼ぶ
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Infof(c, "Hello, world!")
	w.Write([]byte("success!!"))
}

func warmup(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	//log.Infof(c, "Warm Up!")
	fmt.Println("Warm Up!")
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
