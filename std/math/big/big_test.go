package big_test

import (
	"math/big"
	"testing"
)

func TestInt(t *testing.T) {
	a := big.NewInt(1)
	b := big.NewInt(1)
	temp := new(big.Int)
	t.Log(temp, a, b)
	a.Lsh(a, 10)
	b.Lsh(b, 10)
	t.Log(temp, a, b)
	temp.And(a, b)
	// a.And(a, b)
	t.Log(temp, a, b)

	// temp := big.NewInt(0)
	// curr := big.NewInt(1)

	// acc := big.NewInt(1)
	// curr = temp.Lsh(curr, uint(10))
	// t.Log(curr)

	// // t.Log( != big.NewInt(0))
	// temp.And(acc, curr)

	// t.Log(curr)
}
