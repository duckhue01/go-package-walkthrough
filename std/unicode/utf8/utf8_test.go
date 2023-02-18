package utf8_test

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestUTF8(t *testing.T) {
	r := '世'
	buf := make([]byte, 3)

	n := utf8.EncodeRune(buf, r)

	fmt.Println(buf)
	fmt.Println(n)

	res, _ := utf8.DecodeRune(buf)

	fmt.Printf("res: %v\n", string(res))
}
