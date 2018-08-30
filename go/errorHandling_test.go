package _go

import (
	"io"
	"os"
	"testing"
	"log"
	"errors"
)

func Test_Robさんのパターンを試す(t *testing.T) {
	// refs. http://jxck.hatenablog.com/entry/golang-error-handling-lesson-by-rob-pike

	buf := []byte("あろは")

	ew := &errWriter{
		w: os.Stdout, // ファイルではなくstdoutに出力する
	}

	ew.Write(buf)
	ew.Write(buf)

	if ew.Err() != nil {
		return // error
	}
	return // success

}

type errWriter struct {
	w   io.Writer
	err error
}

func (e *errWriter) Write(p []byte) {
	if e.err != nil {
		return
	}
	_, e.err = e.w.Write(p)
}

func (e *errWriter) Err() error {
	return e.err
}

// ---

func A(ctx MyContext) error {
	// なんかやる
	log.Print(ctx.Name)
	return nil // success
}

func B(ctx MyContext) error {
	// なんかやる
	return errors.New("失敗した") // error!!!
}

func C(ctx MyContext) error {
	// なんかやる
	log.Print(ctx.Connection)
	return nil // success
}

func Test_Robさんのパターンで別々の関数を呼ぶパターンを作ってみる(t *testing.T) {

	baseCtx := MyContext{
		Name:"ore",
		Connection:"繋がってる",
	}

	// contextAに処理Aを持たせる
	aCtx := baseCtx
	aCtx.f = A

	// contextBに処理Bを持たせる
	bCtx := baseCtx
	bCtx.f = B

	// contextに処理Cを持たせる
	cCtx := baseCtx
	cCtx.f = C

	// executorを作る。（エラーを保持している）
	ex := errExecutor{
		ctx: aCtx,
	}

	ex.Execute(aCtx) // Aを実行 成功してログが出る
	ex.Execute(bCtx) // Bを実行 errorが発生する
	ex.Execute(cCtx) // Cを実行 errorがBで発生しているので実行されない

	if ex.err != nil {
		t.Log(ex.err.Error())
	}
}

type errExecutor struct {
	err error
	ctx MyContext
}

type MyContext struct {
	Name string
	Connection string // dbコネクションとかのつもり
	f func(context MyContext) error
}

func (x *errExecutor) Execute(context MyContext) {
	if x.err != nil {
		return
	}
	x.err = context.f(context) // ここちょっときもい
}

func(x *errExecutor) Err() error {
	return x.err
}





