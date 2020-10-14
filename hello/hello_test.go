package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Allan")
	want := "Hello, Allan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
