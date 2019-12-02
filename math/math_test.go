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
	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{10.0, 20.0}
		got := rect.Area()
		want := 200.0

		if got != want {
			t.Errorf("got %.2f want %2.f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if want != got {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
}
