package pointer

import "testing"

func TestGetLastVal(t *testing.T) {
	l1 := GenList([]int{2, 4, 3})
	//l2 := GenList([]int{5, 6, 4})

	got := GetLastVal(l1)
	want := 3

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
