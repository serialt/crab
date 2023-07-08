package crab

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func AESEncryptCBC(plaintext, key []byte) (cipherText []byte, err error) {
	block, _ := aes.NewCipher(key)
	data := pkcs7Padding(plaintext, block.BlockSize())

	cipherText = make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], data)
	return

}

func AESEncryptCBCBase64(plaintext, key string) (cipherText string, err error) {
	data, err := AESEncryptCBC([]byte(plaintext), []byte(key))
	cipherText = base64.StdEncoding.EncodeToString(data)
	return
}

func AESDecryptCBC(cipherText, key []byte) (plaintext []byte, err error) {

	block, _ := aes.NewCipher(key)
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		err = errors.New("cipherText is not a multiple of the block size")
		return
	}

	plaintext = make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, cipherText)

	plaintext = pkcs7UnPadding(plaintext)
	return

}

func AESDecryptCBCBase64(cipherText, key string) (plaintext string, err error) {
	_data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return
	}
	data, err := AESDecryptCBC(_data, []byte(key))
	return string(data), err

}

func AESEncryptCTR(plaintext, key []byte) (cipherText []byte, err error) {
	block, _ := aes.NewCipher(key)
	data := pkcs7Padding(plaintext, block.BlockSize())

	cipherText = make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	mode := cipher.NewCTR(block, iv)
	mode.XORKeyStream(cipherText[aes.BlockSize:], data)
	return

}

func AESDecryptCTR(cipherText, key []byte) (plaintext []byte, err error) {

	block, _ := aes.NewCipher(key)
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		err = errors.New("cipherText is not a multiple of the block size")
		return
	}

	plaintext = make([]byte, len(cipherText))
	mode := cipher.NewCTR(block, iv)
	mode.XORKeyStream(plaintext, cipherText)
	plaintext = pkcs7UnPadding(plaintext)

	return
}

func AESEncryptOFB(plaintext, key []byte) (cipherText []byte, err error) {
	block, _ := aes.NewCipher(key)
	data := pkcs7Padding(plaintext, block.BlockSize())

	cipherText = make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(cipherText[aes.BlockSize:], data)
	return

}

func AESDecryptOFB(cipherText, key []byte) (plaintext []byte, err error) {

	block, _ := aes.NewCipher(key)
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		err = errors.New("cipherText is not a multiple of the block size")
		return
	}

	plaintext = make([]byte, len(cipherText))
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(plaintext, cipherText)
	plaintext = pkcs7UnPadding(plaintext)

	return
}

func AESEncryptCFB(plaintext, key []byte) (cipherText []byte, err error) {
	block, _ := aes.NewCipher(key)
	data := pkcs7Padding(plaintext, block.BlockSize())

	cipherText = make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	mode := cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(cipherText[aes.BlockSize:], data)
	return

}

func AESDecryptCFB(cipherText, key []byte) (plaintext []byte, err error) {

	block, _ := aes.NewCipher(key)
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		err = errors.New("cipherText is not a multiple of the block size")
		return
	}

	plaintext = make([]byte, len(cipherText))
	mode := cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(plaintext, cipherText)
	plaintext = pkcs7UnPadding(plaintext)

	return
}

func AESEncryptGCM(plaintext, key []byte) (cipherText []byte, err error) {
	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	cipherText = aesgcm.Seal(nonce, nonce, plaintext, nil)
	return

}

func AESEncryptGCMBase64(plaintext, key string) (cipherText string, err error) {
	data, err := AESEncryptGCM([]byte(plaintext), []byte(key))
	cipherText = base64.StdEncoding.EncodeToString(data)
	return
}

func AESDecryptGCM(cipherText, key []byte) (plaintext []byte, err error) {

	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)
	nonce, cipherText := cipherText[:aesgcm.NonceSize()], cipherText[aesgcm.NonceSize():]
	plaintext, err = aesgcm.Open(nil, nonce, cipherText, nil)
	return
}

func AESDecryptGCMBase64(cipherText, key string) (plaintext string, err error) {
	_data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return
	}
	data, err := AESDecryptGCM(_data, []byte(key))
	return string(data), err

}

func pkcs7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func pkcs7UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}
