//+build wireinject

package main

import "github.com/google/wire"

func FireWire(i int) BType {
	wire.Build(
		A,
		B,
	)
	return ""
}
