package go_xerrors

import (
	"golang.org/x/xerrors"
	"testing"
)

// https://qiita.com/sonatard/items/9c9faf79ac03c20f4ae1

func Test_エラー生成(t *testing.T) {

	err := xerrors.New("エラーだよ")

	t.Logf("%s", err) // 文言のみ表示

	t.Logf("%+v", err) // ファイル名、メソッド名が出る

}

func Test_エラーから新しいエラーを生成(t *testing.T){

	papaErr := xerrors.New("エラー父")

	// 親子一つで出す
	err := xerrors.Errorf("エラー子 -> %v", papaErr)
	t.Logf("%+v", err)

	// 親子分けて出す セミコロン、スペースで分ける
	err2 := xerrors.Errorf("エラー子 : %v", papaErr)
	t.Logf("%+v", err2)

}

func Test_エラーをラップ(t *testing.T) {

	papaErr := xerrors.New("エラー父")

	// ラップは `: %w` でラップする必要がある
	err := xerrors.Errorf("エラー子です : %w", papaErr)

	t.Logf("%+v", err)

}

func Test_エラーをアンラップ(t *testing.T) {

	papaErr := xerrors.New("エラー父")
	err := xerrors.Errorf("エラー子です : %w", papaErr)

	t.Logf("%+v", xerrors.Unwrap(err))

}

func Test_エラーの同一性チェック(t *testing.T) {

	papaErr := xerrors.New("エラー父")
	t.Logf("%v", xerrors.Is(papaErr, papaErr))
	t.Logf("%v", papaErr == papaErr)

}

func Test_エラーの同一性チェック_wrap(t *testing.T) {

	papaErr := xerrors.New("エラー父")
	err := xerrors.Errorf("エラー子です : %w", papaErr)

	t.Logf("%v", xerrors.Is(err, papaErr))
	t.Logf("%v", err == papaErr)

}

func Test_エラーの同一性チェック_wrap複数回(t *testing.T) {

	papaErr := xerrors.New("エラー父")
	sonErr := xerrors.Errorf("エラー子です : %w", papaErr)
	err := xerrors.Errorf("エラー孫です : %w", sonErr)

	t.Logf("%v", xerrors.Is(sonErr, papaErr))
	t.Logf("%v", xerrors.Is(err, papaErr))
	t.Logf("%v", xerrors.Is(err, sonErr))

}
