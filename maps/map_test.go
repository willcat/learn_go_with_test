package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"key": "this is just test value"}
	got := Search(dictionary, "key")
	want := "this is just test value"

	assertString(t, got, want)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
