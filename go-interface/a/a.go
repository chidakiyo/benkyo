package a

import "github.com/chidakiyo/benkyo/go-interface/b"

type A interface {
	A(string)
	AA(int)
}

// 実装されているかチェックは利用する側
var _ A = (*b.B)(nil)

// A型を受け取る
func AAA(v A) {
	// 以下の2つが呼べる
	v.A("") //
	v.AA(0) //
}

// 例えば違うインターフェースでも良い
type A_dash interface {
	AA(int)
}

func AAAA(v A_dash) {
	// 以下だけ呼べる
	v.AA(1)
}