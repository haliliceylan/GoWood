package asymmetric

import (
	"crypto/rand"
	CryptoRSA "crypto/rsa"
	"testing"
)

func TestRSAEncryptDecrypt(t *testing.T) {
	rsa := &RSA{}

	privateKey, _ := CryptoRSA.GenerateKey(rand.Reader, 2048)

	rsa.privateKey = privateKey

	clearText := "test"

	cipherText := rsa.Encrypt([]byte(clearText))

	DecryptedClearText := rsa.Decrypt(cipherText)

	if string(DecryptedClearText) != clearText {
		t.Error("RSA Encryption Decrption is incorrect")
	}
}

func TestRSASignVerify(t *testing.T) {
	rsa := &RSA{}

	privateKey, _ := CryptoRSA.GenerateKey(rand.Reader, 2048)

	rsa.privateKey = privateKey

	clearText := "test"

	signature := rsa.Sign([]byte(clearText))

	verify := rsa.Verify(signature, []byte(clearText))

	if verify {
		t.Error("RSA Sign Verify is incorrect")
	}
}
