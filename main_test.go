package main

import "testing"

func Testdaysecond(t *testing.T) {
	s := daySecond(1, 1, 1)
	if s != 0 {
		t.Errorf("Second was incorrect")
	}
}
