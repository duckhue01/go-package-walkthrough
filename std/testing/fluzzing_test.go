package testing_test

import (
	"errors"
	"testing"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("Input is not valid UTF-8")
	}
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}

func FuzzReverse(f *testing.F) {
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := Reverse(orig)
		if err != nil {
			t.Skip("Skipped")
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}

		var doubleRev string
		doubleRev, err = Reverse(rev)
		if err != nil {
			t.Skip()
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
	})
}

