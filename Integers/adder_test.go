package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		// using %d here to print an integer rather than a string
		t.Errorf("Expected '%d' but got '%d'", expected, sum)
	}
}
