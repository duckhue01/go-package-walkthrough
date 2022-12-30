package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"testing"
)

// Package hmac implements the Keyed-Hash Message Authentication Code.
// An HMAC is cryptographic hash that use a kay to sign a message.
// The receiver verifies it by recomputing it using the same key.

const (
	key = "duckhue01"
)

type (
	Message struct {
		Msg string
		Tag string
	}
)

func TestHMAC(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mac := hmac.New(sha256.New, []byte(key))
		mac.Write([]byte("duckhue01"))
		t.Log(mac.Sum(nil))
	})

	t.Run("", func(t *testing.T) {
		mac := hmac.New(sha256.New, []byte(key))
		// mac.Write([]byte("duckhue01"))
		t.Log(mac.Sum([]byte("duckhue01")))
	})

}

func sign(msg string) Message {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))

	return Message{
		Msg: msg,
		Tag: string(mac.Sum(nil)),
	}
}

func verify(Message Message) bool {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(Message.Msg))

	return hmac.Equal(mac.Sum((nil)), []byte(Message.Tag))
}

func TestSignAndVerify(t *testing.T) {
	msg := sign("test message")
	t.Log(verify(msg))

}
