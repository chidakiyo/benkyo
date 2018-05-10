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
	// --------
	// WithTimeoutは内部的にWithDeadlineを呼んでいる
	// --------
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

func Test_cancelしてみる(t *testing.T) {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go process(ctx)

	time.Sleep(1 * time.Second)
	cancel() // 1秒経ったらキャンセル

	select {
	case <-ctx.Done():
		t.Log("done:", ctx.Err())
	}

}

func Test_cancelの親子関係を確認してみる(t *testing.T) {

	ctx := context.Background()

	parent, pCancel := context.WithCancel(ctx)
	child, cCancel := context.WithCancel(parent)
	defer cCancel()

	go process(ctx)

	go func() {
		time.Sleep(time.Second)
		pCancel()
		t.Log("parent cancelled")
	}()

	select {
	case <-child.Done():
		t.Log("parent:", parent.Err()) // 親をキャンセルすると
		t.Log("child:", child.Err())   // 親をキャンセルすると子もキャンセル
	}

}

func Test_cancelの親子関係の子だけキャンセルしてみる(t *testing.T) {

	ctx := context.Background()

	parent, pCancel := context.WithCancel(ctx)
	child, cCancel := context.WithCancel(parent)
	defer pCancel()

	go process(ctx)

	go func() {
		time.Sleep(time.Second)
		cCancel()
		t.Log("parent cancelled")
	}()

	select {
	case <-child.Done():
		t.Log("parent:", parent.Err()) // 親はキャンセルされない errorはnil
		t.Log("child:", child.Err())   // 子はキャンセル
	}

}
