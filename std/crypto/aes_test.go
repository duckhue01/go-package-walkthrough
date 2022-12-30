package crypto

import (
	"crypto/aes"
	"testing"
)

func TestAES(t *testing.T) {
	bl, err := aes.NewCipher([]byte("asdasd"))
	if err != nil {
	}

	t.Log(bl.BlockSize())
}
