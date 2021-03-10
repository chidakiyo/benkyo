package main

import (
	"context"
	"go.uber.org/ratelimit"
	"sync"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	rl := ratelimit.New(1) // per second

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		t.Log(i, now.Sub(prev))
		prev = now
		time.Sleep(10 * time.Second)
	}
}

func Test2(t *testing.T) {
	rl := ratelimit.New(1) // per second

	var wg sync.WaitGroup
	prev := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			now := rl.Take()
			t.Log(i, now.Sub(prev))
			prev = now
			time.Sleep(10 * time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
}

type IN struct {
	I int
}

func Test3(t *testing.T) {
	rl := ratelimit.New(1) // per second

	var wg sync.WaitGroup
	var cn = make(chan IN, 100)
	prev := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	f := func(ctx context.Context) {
		for {
			now := rl.Take()
			c := <-cn
			t.Log(c.I, now.Sub(prev))
			prev = now
			time.Sleep(10 * time.Second)
			wg.Done()
		}
	}

	for i := 0; i < 8; i++ {
		go f(ctx)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		cn <- IN{i}
	}
	t.Log("push fin.")
	wg.Wait()
	t.Log("fin.")
}
