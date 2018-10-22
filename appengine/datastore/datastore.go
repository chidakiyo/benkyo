package datastore

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	l "log"
	"net/http"
)

// Datatoreに書き込む処理
// go1.9のコードのまま go111 に移行できるか
func DsWrite(w http.ResponseWriter, r *http.Request) {

	l.Printf("START Datastore Write\n")

	// 従来どおりのappengineのcontextを作成する
	ctx := appengine.NewContext(r)

	defer func() {
		if e := recover(); e != nil {
			log.Warningf(ctx, "Error %#v", e)
		}
	}()

	// key作成
	k := datastore.NewKey(ctx, "Entity1", "first", 0, nil)

	// 登録するデータ
	e := Entity1{
		Name:   "first",
		Number: 1,
	}

	// datastoreに書き込み
	if _, err := datastore.Put(ctx, k, e); err != nil {
		// 従来のログ
		log.Errorf(ctx, "datastore put error %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "datastore put success %#v", k)
}

// Datastoreから読み出す処理
// go1.9のコードのまま go111 に移行できるか
func DsRead(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("START Datastore Read\n")

	// 従来どおりのappengineのcontextを作成する
	ctx := appengine.NewContext(r)

	defer func() {
		if e := recover(); e != nil {
			log.Warningf(ctx, "Error %#v", e)
		}
	}()

	// key作成
	k := datastore.NewKey(ctx, "Entity1", "first", 0, nil)

	// 取得するデータ
	e := new(Entity1)

	// datastoreから読み出し
	if err := datastore.Get(ctx, k, e); err != nil {
		// 従来のログ
		log.Errorf(ctx, "datastore get error %s", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "datastore get success %#v", e)
}

type Entity1 struct {
	Number int64
	Name   string
}
