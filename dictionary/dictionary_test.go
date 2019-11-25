// Package main provides ...
package main

import "testing"

func TestSearc(t *testing.T) {
	dictionary := Dictionary{"test": "just a test"}

	got := dictionary.Search("test")
	want := "just a test"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}

}
