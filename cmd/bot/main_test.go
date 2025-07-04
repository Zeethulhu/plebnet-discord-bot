package main

import "testing"

func TestMainDummy(t *testing.T) {
	if false != false {
		t.Error("This should never fail")
	}
}
