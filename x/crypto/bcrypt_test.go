package crypto

import (
	"testing"
	"golang.org/x/crypto/bcrypt"
)

func TestBCrypt(t *testing.T) {
	h, _ := bcrypt.GenerateFromPassword([]byte("duckhue01"), bcrypt.DefaultCost)
	t.Logf("%x", h)
	if bcrypt.CompareHashAndPassword(h, []byte("duckhue01")) != nil {
		t.Log("false")
	}
	t.Log("true")

}
