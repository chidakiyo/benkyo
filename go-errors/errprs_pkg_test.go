package main_test

import (
	"github.com/pkg/errors"
	"testing"
)

func Test_error_wrap(t *testing.T) {
	e := errors.New("ベース")
	e = errors.Wrap(e, "階層 1")
	e = errors.Wrapf(e, "階層 %d", 2)

	t.Logf("string : %s\n", e.Error())
	t.Logf("string : %s\n", e)
	t.Logf("struct : %+v\n", e)
}

