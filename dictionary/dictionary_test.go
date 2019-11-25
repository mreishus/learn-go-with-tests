// Package main provides ...
package main

import "testing"

func TestSearc(t *testing.T) {
	dictionary := Dictionary{"test": "just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		want := ErrNotFound

		assertError(t, got, want)
	})
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected to get an error")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}

}
