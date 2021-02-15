package main

import "testing"

func Test_スライスがnilのときの1個めの要素(t *testing.T) {

	t.Run("0番地指定", func(t *testing.T) {
		var x []string

		target := x[0]
		t.Logf("%v", target)

	})

	t.Run("切り取る", func(t *testing.T) {
		var x = []string{}

		target := x[:1]
		t.Logf("%v", target)
	})

	// 結果どっちも同じふるまい

}
