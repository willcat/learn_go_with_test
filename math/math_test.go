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
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f want %2.f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{10.0, 20.0}
		want := 200.0

		checkArea(t, rect, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})
}
