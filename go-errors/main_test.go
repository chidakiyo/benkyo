package main_test

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func Test_error_wrap(t *testing.T) {
	e := errors.New("error!!")
	e = errors.Wrap(e, "message 1")
	e = errors.Wrapf(e, "message %d", 2)

	fmt.Printf("string : %s\n", e)
	fmt.Printf("struct : %+v\n", e)
}

