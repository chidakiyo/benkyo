package lib

import "testing"

func Test_gen_byte(t *testing.T) {
	b := genBytes()
	t.Logf("%v", b)
}
