package main

import (
	"cloud.google.com/go/compute/metadata"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
	"os"
)

func main() {
	route := gin.Default()
	http.Handle("/", route)

	route.GET("/", root)
	route.GET("/header", headers)
	route.GET("/env", env)

	// GAE上のmetadataを取得する
	route.GET("/gcp", gcpMeta)

	appengine.Main() // Listen
}

func root(g *gin.Context) {

	fmt.Println("ACCESS!!!")

	g.String(http.StatusOK, "ok")
}

func headers(g *gin.Context) {

	for k, v := range g.Request.Header {
		fmt.Printf("%s, %v\n", k, v)
	}

	g.String(http.StatusOK, "ok")
}

func env(g *gin.Context) {

	for _, v := range os.Environ() {
		fmt.Printf("%v\n", v)
	}

	g.String(http.StatusOK, "ok")
}

func gcpMeta(g *gin.Context) {

	ip, _ := metadata.ExternalIP()
	fmt.Printf("IP: %s\n", ip)

	host, _ := metadata.Hostname()
	fmt.Printf("HOST: %s\n", host)

	att, _ := metadata.InstanceAttributes()
	fmt.Printf("Atts: %+v\n", att)

	iid, _ := metadata.InstanceID()
	fmt.Printf("IID: %s\n", iid)

	name, _ := metadata.InstanceName()
	fmt.Printf("NAME: %s\n", name)

	iTag, _ := metadata.InstanceTags()
	fmt.Printf("iTag: %+v\n", iTag)

	iip, _ := metadata.InternalIP()
	fmt.Printf("iIP: %s\n", iip)

	npid, _ := metadata.NumericProjectID()
	fmt.Printf("numPID: %s\n", npid)

	pid, _ := metadata.ProjectID()
	fmt.Printf("PID: %s\n", pid)

	patt, _ := metadata.ProjectAttributes()
	fmt.Printf("pATT: %+v\n", patt)

	g.String(http.StatusOK, "end")

}
