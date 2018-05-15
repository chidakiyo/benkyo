package go_sandbox

import (
	"testing"
	"github.com/pkg/errors"
)

func Test_recoverでエラーログを出す(t *testing.T) {

	func() {

		defer func() {
			if err := recover(); err != nil {
				t.Logf("パニックを補足した %s \n", err)
				// なんか処理
			}

		}()

		panic("パニパニパニック！")

	} ()

}

func Test_recoverの戻りがinterfaceなのはpanicの引数を受け取るから(t *testing.T) {

	func() {

		defer func() {
			if err := recover(); err != nil {
				t.Logf("パニックを補足した %s \n", err.(error).Error()) // recoverの戻りがerror型
				// なんか処理
			}

		}()

		panic(errors.New("パニパニパニック！")) // panicにerror型で渡すから ↑

	} ()

}
