package math

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 20.0}
	got := Perimeter(rect)
	want := float64(60)

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rect := Rectangle{10.0, 20.0}
	got := Area(rect)
	want := 200.0

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}
