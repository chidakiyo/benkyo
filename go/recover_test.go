package _go

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

func Test_recoverで戻り値を返さない場合に何が返るか_bool(t *testing.T) {

	result := func() bool {

		defer func() {
			if err := recover(); err != nil {
				t.Logf("パニックを補足した %s \n", err)
				// なんか処理
			}

		}()

		panic("パニパニパニック！")

		return true

	} ()

	t.Log(result) // falseが返る
}

func Test_recoverで戻り値を返さない場合に何が返るか_string(t *testing.T) {

	result := func() string {

		defer func() {
			if err := recover(); err != nil {
				t.Logf("パニックを補足した %s \n", err)
				// なんか処理
			}

		}()

		panic("パニパニパニック！")

		return "成功！"

	} ()

	t.Log(result) // ""が返る
}

func Test_recoverで戻り値を返さない場合に何が返るか_int(t *testing.T) {

	result := func() int64 {

		defer func() {
			if err := recover(); err != nil {
				t.Logf("パニックを補足した %s \n", err)
				// なんか処理
			}

		}()

		panic("パニパニパニック！")

		return 1000

	} ()

	t.Log(result) // 0が返る
}

func Test_recoverで戻り値を返す場合(t *testing.T) {

	result := func() (e string)  {

		defer func() {
			if err := recover(); err != nil {
				t.Logf("パニックを補足した %s \n", err)
				// なんか処理
				e = "失敗" // 戻り値 e を上書きする
			}

		}()

		panic("パニパニパニック！")

		return "成功"

	} ()

	t.Log(result) // 失敗 が返る
}