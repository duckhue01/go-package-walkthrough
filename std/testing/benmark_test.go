package testing_test

import (
	"fmt"
	"testing"
)

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

// 1000000000	         0.0004851 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPointer(b *testing.B) {
	type A struct {
		name string
		age  int
	}

	fun := func(a *A) {
		fmt.Print(a.age, a.name)
	}

	AInstant := &A{
		name: "duckhue01",
		age:  100,
	}

	for i := 0; i < 100; i++ {
		fun(AInstant)
	}

}

// 1000000000	         0.0004288 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPointer01(b *testing.B) {
	type A struct {
		name        string
		age         int
		siblings    uint16
		passportNum uint64
	}

	fun := func(a A) {
		fmt.Print(a.age, a.name)
	}

	AInstant := A{
		name: "duckhue01",
		age:  100,
	}

	for i := 0; i < 100; i++ {
		fun(AInstant)
	}

}
