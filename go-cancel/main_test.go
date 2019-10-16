package main

import (
	"context"
	"testing"
	"time"
)

func slowFunction(sec time.Duration) {
	time.Sleep(sec)
}

func Test_Cancelでキャンセル(t *testing.T) {

	c1 := context.Background()
	_, cancel := context.WithCancel(c1)
	defer func() {
		t.Log("defer")
		cancel()
	}()

	go func() {
		t.Log("goroutine start")
		<-time.After(3 * time.Second)
		cancel()
		t.Log("goroutine cancel.")
	}()

	t.Log("slowfunc start")
	slowFunction(5 * time.Second)
	t.Log("slowfunc end")

	t.Logf("finish")
}







