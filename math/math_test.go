package math

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10, 20)
	want := float64(60)

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 20.0)
	want := 200.0

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}
