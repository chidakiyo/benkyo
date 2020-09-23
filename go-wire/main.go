package main

import "fmt"

func main() {
	result := FireWire(10)
	fmt.Println(result)
}

type AType string

func A(a int) AType {
	return AType(fmt.Sprintf("%d", a))
}

type BType string

func B(b AType) BType {
	return BType(fmt.Sprintf("[%s]", b))
}



