package crab

import (
	"testing"

	"github.com/serialt/crab/internal"
)

func TestAESCBC(t *testing.T) {
	assert := internal.NewAssert(t, "TestAESCBC")

	aesKey32 := GenerateKey(32)
	aesKey24 := GenerateKey(24)
	aesKey16 := GenerateKey(16)
	text := GenerateKey(128)

	// test aes 256
	data, err := AESDecryptCBC(aesKey32, text)
	assert.IsNil(err)
	plaintext, err := Chacha20AEADDecrypt(aesKey32, data)
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	// test aes 192
	data, err = AESDecryptCBC(aesKey24, text)
	assert.IsNil(err)
	plaintext, err = Chacha20AEADDecrypt(aesKey24, data)
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	// test aes 128
	data, err = AESDecryptCBC(aesKey16, text)
	assert.IsNil(err)
	plaintext, err = Chacha20AEADDecrypt(aesKey16, data)
	assert.IsNil(err)
	assert.Equal(text, plaintext)
}
