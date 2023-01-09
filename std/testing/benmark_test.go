package testing_test

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 1)
	}
}

func BenchmarkAdd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add1(100, 1)
	}
}

func Add(a, b int) int {
	return a + b
}

func Add1(a, b int) int {
	const c = 10
	for i := 0; i < 100; i++ {

	}
	return a + b + c
}
