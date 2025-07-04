package main

import "testing"

func TestStartupMessage(t *testing.T) {
	got := getStartupMessage()
	want := "Bot started."

	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
