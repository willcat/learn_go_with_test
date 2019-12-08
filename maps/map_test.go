package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"key": "this is just test value"}
	got := dictionary.Search("key")
	want := "this is just test value"

	assertString(t, got, want)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
