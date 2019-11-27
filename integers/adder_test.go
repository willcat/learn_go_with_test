package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("testing adder", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d' ", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(3, 5)
	fmt.Println(sum)
	// Output: 6
}
