// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func FireWire(i int) BType {
	aType := A(i)
	bType := B(aType)
	return bType
}
