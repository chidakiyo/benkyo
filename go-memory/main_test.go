package main

import "testing"

func BenchmarkLoopp(b *testing.B) {
	res := make([]int, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = append(res, len(loopp()))
	}
}

func BenchmarkLoopnp(b *testing.B) {
	res := make([]int, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = append(res, len(loopnp()))
	}
}
