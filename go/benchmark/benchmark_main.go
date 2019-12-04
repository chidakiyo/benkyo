package main

import (
	"fmt"
	"testing"
	"time"
)

func main() {
	result := testing.Benchmark(func(b *testing.B) {
		b.ResetTimer()
		for i := 0 ; i <= b.N ; i++{
			time.Sleep(1 * time.Millisecond)
		}
	})
	fmt.Printf("%v", result)
}
