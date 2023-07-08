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
	data, err := AESEncryptCBC(text, aesKey32)
	t.Logf("data: %s, key: %v", string(text), string(aesKey32))
	assert.IsNil(err)
	plaintext, err := AESDecryptCBC(data, aesKey32)
	t.Logf("plaintext: %s", string(plaintext))
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	//aes 256 base64
	_data, err := AESEncryptCBCBase64(string(text), string(aesKey32))
	t.Logf("data: %s, key: %v", string(text), string(aesKey32))
	assert.IsNil(err)
	_plaintext, err := AESDecryptCBCBase64(_data, string(aesKey32))
	t.Logf("plaintext: %s", _plaintext)
	assert.IsNil(err)
	assert.Equal(string(text), _plaintext)

	// test aes 192
	data, err = AESEncryptCBC(text, aesKey24)
	assert.IsNil(err)
	plaintext, err = AESDecryptCBC(data, aesKey24)
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	//aes 192 base64
	_data, err = AESEncryptCBCBase64(string(text), string(aesKey24))
	t.Logf("data: %s, key: %v", string(text), string(aesKey24))
	assert.IsNil(err)
	_plaintext, err = AESDecryptCBCBase64(_data, string(aesKey24))
	t.Logf("plaintext: %s", _plaintext)
	assert.IsNil(err)
	assert.Equal(string(text), _plaintext)

	// test aes 128
	data, err = AESEncryptCBC(text, aesKey16)
	assert.IsNil(err)
	plaintext, err = AESDecryptCBC(data, aesKey16)
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	//aes 128 base64
	_data, err = AESEncryptCBCBase64(string(text), string(aesKey16))
	t.Logf("data: %s, key: %v", string(text), string(aesKey16))
	assert.IsNil(err)
	_plaintext, err = AESDecryptCBCBase64(_data, string(aesKey16))
	t.Logf("plaintext: %s", _plaintext)
	assert.IsNil(err)
	assert.Equal(string(text), _plaintext)

}

func TestAESCFB(t *testing.T) {
	assert := internal.NewAssert(t, "TestAESCFB")

	aesKey32 := GenerateKey(32)
	aesKey24 := GenerateKey(24)
	aesKey16 := GenerateKey(16)
	text := GenerateKey(128)

	// test aes 256
	data, err := AESEncryptCFB(text, aesKey32)
	t.Logf("data: %s, key: %v", string(text), string(aesKey32))
	assert.IsNil(err)
	plaintext, err := AESDecryptCFB(data, aesKey32)
	t.Logf("plaintext: %s", string(plaintext))
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	// test aes 192
	data, err = AESEncryptCFB(text, aesKey24)
	assert.IsNil(err)
	plaintext, err = AESDecryptCFB(data, aesKey24)
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	// test aes 128
	data, err = AESEncryptCBC(text, aesKey16)
	assert.IsNil(err)
	plaintext, err = AESDecryptCBC(data, aesKey16)
	assert.IsNil(err)
	assert.Equal(text, plaintext)
}

func TestAESCTR(t *testing.T) {
	assert := internal.NewAssert(t, "TestAESCTR")

	aesKey32 := GenerateKey(32)
	aesKey24 := GenerateKey(24)
	aesKey16 := GenerateKey(16)
	text := GenerateKey(128)

	// test aes 256
	data, err := AESEncryptCTR(text, aesKey32)
	t.Logf("data: %s, key: %v", string(text), string(aesKey32))
	assert.IsNil(err)
	plaintext, err := AESDecryptCTR(data, aesKey32)
	t.Logf("plaintext: %s", string(plaintext))
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	// test aes 192
	data, err = AESEncryptCTR(text, aesKey24)
	assert.IsNil(err)
	plaintext, err = AESDecryptCTR(data, aesKey24)
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	// test aes 128
	data, err = AESEncryptCTR(text, aesKey16)
	assert.IsNil(err)
	plaintext, err = AESDecryptCTR(data, aesKey16)
	assert.IsNil(err)
	assert.Equal(text, plaintext)
}

func TestAESGCM(t *testing.T) {
	assert := internal.NewAssert(t, "TestAESGCM")

	aesKey32 := GenerateKey(32)
	aesKey24 := GenerateKey(24)
	aesKey16 := GenerateKey(16)
	text := GenerateKey(128)

	// test aes 256
	data, err := AESEncryptGCM(text, aesKey32)
	t.Logf("data: %s, key: %v", string(text), string(aesKey32))
	assert.IsNil(err)
	plaintext, err := AESDecryptGCM(data, aesKey32)
	t.Logf("plaintext: %s", string(plaintext))
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	// test aes 192
	data, err = AESEncryptGCM(text, aesKey24)
	assert.IsNil(err)
	plaintext, err = AESDecryptGCM(data, aesKey24)
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	// test aes 128
	data, err = AESEncryptGCM(text, aesKey16)
	assert.IsNil(err)
	plaintext, err = AESDecryptGCM(data, aesKey16)
	assert.IsNil(err)
	assert.Equal(text, plaintext)
}

func TestAESOFB(t *testing.T) {
	assert := internal.NewAssert(t, "TestAESOFB")

	aesKey32 := GenerateKey(32)
	aesKey24 := GenerateKey(24)
	aesKey16 := GenerateKey(16)
	text := GenerateKey(128)

	// test aes 256
	data, err := AESEncryptOFB(text, aesKey32)
	t.Logf("data: %s, key: %v", string(text), string(aesKey32))
	assert.IsNil(err)
	plaintext, err := AESDecryptOFB(data, aesKey32)
	t.Logf("plaintext: %s", string(plaintext))
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	// test aes 192
	data, err = AESEncryptOFB(text, aesKey24)
	assert.IsNil(err)
	plaintext, err = AESDecryptOFB(data, aesKey24)
	assert.IsNil(err)
	assert.Equal(text, plaintext)

	// test aes 128
	data, err = AESEncryptOFB(text, aesKey16)
	assert.IsNil(err)
	plaintext, err = AESDecryptOFB(data, aesKey16)
	assert.IsNil(err)
	assert.Equal(text, plaintext)
}
