package crab

import (
	"testing"

	"github.com/serialt/crab/internal"
)

func TestRSAAOEP(t *testing.T) {
	assert := internal.NewAssert(t, "TestRSA")
	var keyLen = []int{2048, 4096, 8192}
	for _, k := range keyLen {
		priKey, pubKey, err := GenerateRSAKey(k)
		t.Logf("priKey:\n %s", string(priKey))
		t.Logf("pubKey:\n %s", string(pubKey))
		assert.IsNil(err)
		text := GenerateKey(128)
		data, err := RSAEncryptOAEP(text, pubKey)
		assert.IsNil(err)
		plaintext, err := RSADecryptOAEP(data, priKey)
		assert.IsNil(err)
		assert.Equal(text, plaintext)
	}

}

func TestRSAAOEPPwd(t *testing.T) {
	assert := internal.NewAssert(t, "TestRSA")
	pass := []byte("sugar")
	var keyLen = []int{2048, 4096, 8192}
	for _, k := range keyLen {
		priKey, pubKey, err := GenerateRSAKeyWithPwd(pass, k)
		t.Logf("priKey:\n %s", string(priKey))
		t.Logf("pubKey:\n %s", string(pubKey))
		assert.IsNil(err)
		text := GenerateKey(128)
		data, err := RSAEncryptOAEP(text, pubKey)
		assert.IsNil(err)
		plaintext, err := RSADecryptOAEPPwd(data, priKey, pass)
		assert.IsNil(err)
		assert.Equal(text, plaintext)
	}

}
