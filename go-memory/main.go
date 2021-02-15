package main

import (
	"fmt"
	"math"
)

func main() {
	x := loopp()
	fmt.Printf("END %d\n", len(x))
}

type X struct {
	String string
	Int    int
}

//const max = math.MaxInt16 * 100
const max = math.MaxInt16

func loopp() []*X {
	x := make([]*X, max)
	for i := 0; i < max; i++ {
		x = append(x, &X{
			String: "TEST",
			Int:    i,
		})
	}
	return x
}

func loopnp() []X {
	x := make([]X, max)
	for i := 0; i < max; i++ {
		x = append(x, X{
			String: "TEST",
			Int:    i,
		})
	}
	return x
}
