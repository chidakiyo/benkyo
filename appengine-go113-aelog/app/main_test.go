package main

import (
	"context"
	log "github.com/DeNA/aelog"
	"testing"
)

func Test_ベースのテスト(t *testing.T) {
	c := context.Background()

	base(c, "hello!")
}

func base(ctx context.Context, t string) {
	log.Infof(ctx, "info message %s", t)
}
