package test

import "testing"

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func Equal[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
