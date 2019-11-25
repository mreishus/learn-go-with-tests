// Package main provides ...
package main

import "testing"

func TestSearc(t *testing.T) {
	dictionary := map[string]string{"test": "just a test"}

	got := Search(dictionary, "test")
	want := "just a test"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}

}
