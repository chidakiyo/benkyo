package b

import "fmt"

// 利用される側はinterfaceを意識しない
// Interfaceがこの実装に入ってこないため、不純物が入らない
type B struct {}

func (B) A(a string){
	fmt.Println(a)
}

func (B) AA(a int){
	fmt.Println(a)
}

var BB = B{}