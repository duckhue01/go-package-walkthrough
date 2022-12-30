package crypto

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/sha3"
)

// Package sha3 implements the SHA-3 fixed-output-length hash functions and the SHAKE variable-output-length hash functions defined by FIPS-202.

// Both types of hash function use the "sponge" construction and the Keccak permutation. For a detailed specification see http://keccak.noekeon.org/

func TestSha3(t *testing.T) {

}

func TestShake256(t *testing.T) {
	buf := []byte("and this is some data to authenticate")
	out := make([]byte, 100)
	hash := sha3.NewShake256()
	hash.Write(buf)
	// Read 32 bytes of output from the hash into h.
	hash.Read(out)
	fmt.Printf("%x\n", out)
}

func TestCShake(t *testing.T) {
	out := make([]byte, 100)
	buf := []byte("hello world")
	hash1 := sha3.NewCShake256([]byte{1}, []byte{1})
	hash1.Write(buf)
	hash1.Read(out)
	fmt.Printf("%x\n", out)

	hash2 := sha3.NewCShake256([]byte{1}, []byte{1})
	hash2.Write(buf)
	hash2.Read(out)
	fmt.Printf("%x\n", out)

}
