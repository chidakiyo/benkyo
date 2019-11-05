package lib

import (
	"errors"
	"fmt"
	"testing"
)

func Test_gen_byte(t *testing.T) {
	b := genBytes(8)
	t.Logf("%v", b)
}

func Test_err_go13(t *testing.T) {

	result1 := func() error {
		return errors.New("new error")
	} ()

	result2 := func() error {
		return nil
	}()

	result3 := func() error {
		return fmt.Errorf("%w", error(nil))
	}()

	if result1 != nil {
		t.Logf("result is error! %s", result1.Error())
	}

	if result2 != nil {
		t.Logf("result2 is error! %s", result2.Error())
	}

	if result3 != nil {
		t.Logf("result3 is error! %s", result3.Error()) // ホントは入らない
	}

}
