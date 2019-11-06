package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Repeat a 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Repeat a 3 times", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Repeat b 4 times", func(t *testing.T) {
		got := Repeat("b", 4)
		want := "bbbb"
		assertCorrectMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
