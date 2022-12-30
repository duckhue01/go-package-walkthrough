package crypto

import (
	"crypto/sha1"
	"testing"
)

// SHA-1 is cryptographically broken and should not be used for secure applications.

// func TestMain(m *testing.M) {
// 	log.Println("Do stuff BEFORE the tests!")
// 	exitVal := m.Run()
// 	log.Println("Do stuff AFTER the tests!")

// 	os.Exit(exitVal)
// }
func TestSha1(t *testing.T) {
	data1 := []byte("1123")
	data2 := []byte("1")
	t.Logf("%t", sha1.Sum(data1) == sha1.Sum(data2))
}
