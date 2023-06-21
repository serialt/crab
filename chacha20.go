package crab

import (
	"math/rand"

	"golang.org/x/crypto/chacha20poly1305"
)

// aeadEncrypt encrypts a message with a one-time key.
func Chacha20AEADEncrypt(plaintext, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, chacha20poly1305.NonceSize)
	return aead.Seal(nil, nonce, plaintext, nil), nil
}

func Chacha20AEADDecrypt(ciphertext, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, chacha20poly1305.NonceSize)
	return aead.Open(nil, nonce, ciphertext, nil)
}

// GenerateKey aes bit: 16/24/32  chacha: 32
func GenerateKey(bit int) []byte {
	return []byte(randStr(bit))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
