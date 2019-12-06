package _go

import (
	"testing"
	"time"
)

func Benchmark_Sample(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 遅い処理
		time.Sleep(10 * time.Millisecond)
	}
}
