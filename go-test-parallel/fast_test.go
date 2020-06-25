package go_test_parallel

import (
	"testing"
	"time"
)

func Test_fast1(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)

}

func Test_fast2(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func Test_fast3(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

