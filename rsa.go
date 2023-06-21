package crab

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// GenerateRSAKey 创建RSA 常用位数 1024 2048 4096
func GenerateRSAKey(bits int) (priKey, pubKey []byte, err error) {
	priWriter := bytes.NewBuffer([]byte{})
	pubWriter := bytes.NewBuffer([]byte{})

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}

	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: X509PrivateKey,
	}

	if err = pem.Encode(priWriter, privateBlock); err != nil {
		return
	}

	publicKey := privateKey.PublicKey
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}

	publicBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: X509PublicKey,
	}

	if err = pem.Encode(pubWriter, publicBlock); err != nil {
		return
	}
	priKey = priWriter.Bytes()
	pubKey = pubWriter.Bytes()
	return
}

// GenerateRSAKeyWithPwd 创建带密码的RSA
func GenerateRSAKeyWithPwd(passwd string, bits int) (priKey, pubKey []byte, err error) {

	priWriter := bytes.NewBuffer([]byte{})
	pubWriter := bytes.NewBuffer([]byte{})

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock, err := x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", x509PrivateKey, []byte(passwd), x509.PEMCipherAES256)
	if err != nil {
		return
	}
	err = pem.Encode(priWriter, privateBlock)
	if err != nil {
		return
	}
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return
	}
	publicBlock := &pem.Block{Type: "RSA PUBLIC KEY", Bytes: X509PublicKey}
	err = pem.Encode(pubWriter, publicBlock)
	if err != nil {
		return
	}
	priKey = priWriter.Bytes()
	pubKey = pubWriter.Bytes()
	return
}

// RSAEncryptOAEP 公钥加密
func RSAEncryptOAEP(plainText, pubCipherKey []byte) (cipherText []byte, err error) {
	block, _ := pem.Decode(pubCipherKey)
	if block == nil {
		err = fmt.Errorf("failed to parse certificate PEM")
		return
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	publickey := pub.(*rsa.PublicKey)
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, publickey, plainText, nil)
}

// RSADecryptOAEP 私钥解密
func RSADecryptOAEP(cipherText, privCipherKey []byte) (plainText []byte, err error) {
	block, _ := pem.Decode(privCipherKey)
	if block == nil {
		err = fmt.Errorf("failed to parse certificate PEM")
		return
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipherText, nil)
}
