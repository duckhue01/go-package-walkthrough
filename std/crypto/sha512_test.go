package crypto

import (
	"crypto/sha512"
	"testing"
)

func TestSha512(t *testing.T) {

	t.Logf("%x", sha512.Sum384([]byte("hello world\n")))
	t.Logf("%x", sha512.Sum512([]byte("hello world\n")))
	t.Logf("%x", sha512.Sum512_224([]byte("hello world\n")))
	t.Logf("%x", sha512.Sum512_256([]byte("hello world\n")))
}
