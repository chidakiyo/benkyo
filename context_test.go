package go_sandbox

import (
	"context"
	"testing"
	"time"
)

func Test_contextに値を入れてみる(t *testing.T) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, "A", 1)
	ctx = context.WithValue(ctx, "B", "あ")

	t.Log(ctx.Value("A")) // 1
	t.Log(ctx.Value("B")) // あ

}

func Test_contextにstruct入れてみる(t *testing.T) {

	ctx := context.Background()

	type A struct {
		x int64
		y string
	}

	a := A{
		1,
		"あ",
	}

	ctx = context.WithValue(ctx, "A", a)

	t.Log(ctx.Value("A")) // {1 あ}

	// ---- appendix ----
	// contextから取得したstructをstructとして扱う

	aa := ctx.Value("A")

	// aa.x  <- interfaceなのでプロパティへアクセスできない

	t.Log(aa.(A).x) // A型で型アサーションして触る

}

func Test_timeoutを使ってみる(t *testing.T) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second) // 1秒後にキャンセル
	defer func() {
		t.Log("defer")
		cancel()
	}()

	go process(ctx)

	select {
	case <-ctx.Done(): // timeoutされるとctxがcloseされてここに入る
		t.Log("done:", ctx.Err())
	}
}

// 重い処理
func process(ctx context.Context) {
	time.Sleep(5 * time.Second)
}
