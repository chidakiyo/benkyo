package main

import (
	"errors"
	"fmt"
	"testing"
)

var (
	aerr = &AErr{}
	berr = &BErr{}
)

func Test_error_wrap(t *testing.T) {

	err := errors.New("ベース")

	w1err := Wrap(err)

	waerr := aerr.Wrap(err)
	wberr := berr.Wrap(err)

	t.Logf("%T", err)
	t.Logf("%T", w1err)

	t.Logf("%T", waerr)
	t.Logf("%T", wberr)

	if errors.As(waerr, aerr) {
		t.Log("a is A")
	} else {
		t.Log("おかしい")
	}

	if errors.As(wberr, berr) {
		t.Log("b is B")
	} else {
		t.Log("おかしい")
	}
}

func Wrap(err error) error {
	errors.New("WrapError")
	return fmt.Errorf("%w", err)
}

type AErr struct {
}

func (e AErr) Wrap(err error) error {
	return fmt.Errorf("%w", err)
}

func (e AErr) Error() string {
	return "A"
}

type BErr struct {
}

func (e BErr) Wrap(err error) error {
	return fmt.Errorf("%w", err)
}

func (e BErr) Error() string {
	return "B"
}

func Unwrap(err error) error {
	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return nil
	}
	return u.Unwrap()
}

func Test_errorがnilのときにswitchできるんだっけ(t *testing.T) {

	// 適当な関数、errorはnilを返すk
	_, err := func() (string, error) {
		return "result", nil
	}()

	switch err {
	case nil:
		t.Log("nilだよ")
	default:
		t.Log("その他だよ")
	}

}
