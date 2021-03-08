package main

import (
	"github.com/chidakiyo/benkyo/go-interface/a"
	b2 "github.com/chidakiyo/benkyo/go-interface/b"
)

func main() {

	// Bのインスタンス生成（インターフェースではなくconcrete）
	b := b2.B{}

	// 呼べる
	a.AAA(b)

	// 呼べる
	a.AAAA(b)
}

