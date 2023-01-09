package quick_test

import (
	"testing"
	"testing/quick"
)

func subtract(a, b int8) int8 {
	return a - b
}

func FuzzSubtract(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int8) {
		got := subtract(a, b)
		t.Log(got)
	})
}

func TestSubtractPropertyFour(t *testing.T) {
	property := func(a, b int8) bool {
		if b > 0 && b < a {
			return subtract(a, b) < a
		}

		return true

	}
	if err := quick.Check(property, nil); err != nil {
		t.Error(err)
	}
}
