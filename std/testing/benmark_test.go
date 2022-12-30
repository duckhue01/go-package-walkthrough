package testing_test

import "testing"

// 229891120	         5.213 ns/op	       0 B/op	       0 allocs/op
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 1)
	}
}

// 23744432	        49.98 ns/op	       0 B/op	       0 allocs/op
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
