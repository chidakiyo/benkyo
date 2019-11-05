package lib

import "testing"

func Test_gen_byte(t *testing.T) {
	b := genBytes(8)
	t.Logf("%v", b)
}
