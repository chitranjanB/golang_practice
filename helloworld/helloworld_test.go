package helloworld

import "testing"

func TestGetHomeDir(t *testing.T) {
	got := GetHomeDir()
	want := "/Users/siva"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
