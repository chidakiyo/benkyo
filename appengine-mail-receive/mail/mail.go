package mail

import (
	"github.com/curious-eyes/jmail"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func Mail(g *gin.Context) {

	r := g.Request
	defer r.Body.Close()

	p := g.Param("address")

	ctx := appengine.NewContext(r)
	log.Infof(ctx, "address %s", p)

	msg, err := jmail.ReadMessage(r.Body)
	if err != nil {
		log.Errorf(ctx, "read fail %s", err)
		return
	}

	log.Infof(ctx, "mail dump body : %#v", msg.Body)
	log.Infof(ctx, "mail dump header : %#v", msg.Header)
	log.Infof(ctx, "mail dump msg : %#v", msg.Message)
	b, _ := msg.DecBody()
	log.Infof(ctx, "mail dump dec_body : %#v", string(b))
	log.Infof(ctx, "mail dump dec_subject : %#v", msg.DecSubject())

	g.String(http.StatusOK, "ok")
}
