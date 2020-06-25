package go_test_parallel

import (
	"testing"
	"time"
)

func Test_slow1(t *testing.T) {
	time.Sleep(3 * time.Second)

}

func Test_slow2(t *testing.T) {
	time.Sleep(1 * time.Second)

}

func Test_slow3(t *testing.T) {
	time.Sleep(2 * time.Second)

}

