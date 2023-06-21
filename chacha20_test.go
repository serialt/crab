package crab

import (
	"testing"

	"github.com/serialt/crab/internal"
)

func TestChachaAEAD(t *testing.T) {
	assert := internal.NewAssert(t, "TestChacha20")

	chacha20Key := GenerateKey(32)

	text := GenerateKey(64)
	data, err := Chacha20AEADEncrypt(text, chacha20Key)
	assert.IsNil(err)

	plaintext, err := Chacha20AEADDecrypt(data, chacha20Key)
	assert.IsNil(err)
	assert.Equal(text, plaintext)

}
