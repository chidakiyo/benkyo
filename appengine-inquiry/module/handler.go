package module

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Path(g *gin.Context){

	c := appengine.NewContext(g.Request)

	p, ok := g.GetQuery("p")
	if !ok {
		g.String(http.StatusBadRequest, "query not found.")
		return
	}

	var ppp []string
	err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		log.Infof(c, "path : %s", path)
		ppp = append(ppp, path)
		return nil
	})

	if err != nil {
		g.String(http.StatusInternalServerError, "cause %s", err.Error())
		return
	}

	g.String(http.StatusOK, "%s", strings.Join(ppp, "\n"))
}