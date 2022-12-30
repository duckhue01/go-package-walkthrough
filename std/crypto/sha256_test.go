package crypto

import (
	"crypto/sha256"
	"testing"
)

const (
	name = "ceddacb1c1fa70240942ee7942dbb7e1232a293ebb3f2af11481e246900422c4"
)

func TestSha256(t *testing.T) {

	hash := sha256.New()
	hash.Write([]byte("duckhue01"))
	t.Logf("%x", string(hash.Sum(nil)))
	// t.Logf("%t", string(hash.Sum(nil)) == name)

}
