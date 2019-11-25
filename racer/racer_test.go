package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("Works", func(t *testing.T) {
		t.Helper()
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("times out after 10 seconds", func(t *testing.T) {
		t.Helper()
		aServer := makeDelayedServer(11 * time.Second)
		bServer := makeDelayedServer(12 * time.Second)

		defer aServer.Close()
		defer bServer.Close()

		_, err := Racer(aServer.URL, bServer.URL)

		if err == nil {
			t.Error("expected a timeout error, but didn't get one")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(delay)
				w.WriteHeader(http.StatusOK)
			}))
}
