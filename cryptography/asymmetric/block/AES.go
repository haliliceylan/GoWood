package block

import (
	cryptoAES "crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

type AES struct {
	publicKey cipher.Block
	mode      cipher.AEAD
	nonce     []byte
}

func (aes *AES) SetPublicKey(key []byte) {
	var err error
	aes.publicKey, err = cryptoAES.NewCipher(key)

	if err != nil {
		log.Fatalf("Error during the Public Key Generating reason: %s", err.Error())
	}
}

func (aes *AES) selectMode() {
	var err error

	aes.mode, err = cipher.NewGCM(aes.publicKey)

	if err != nil {
		log.Fatalf("Error during the Selecting AES Mode reason: %s", err.Error())
	}
}

func (aes *AES) createNoonce() {
	var err error
	aes.nonce = make([]byte, aes.mode.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, aes.nonce); err != nil {
		log.Fatalf("Error during the generating AES Nonce reason: %s", err)
	}
}
func (aes *AES) Encrypt(input []byte) []byte {
	aes.selectMode()
	aes.createNoonce()
	return aes.mode.Seal(aes.nonce, aes.nonce, input, nil)
}

func (aes *AES) Decrypt(ciphertext []byte) []byte {
	aes.selectMode()
	nonceSize := aes.mode.NonceSize()
	if len(ciphertext) < nonceSize {
		log.Fatalf("cipher text is too short expected %v, got %v", nonceSize, len(ciphertext))
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aes.mode.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Fatalf("Error during  the decryption reason: %s", err.Error())
	}

	return plaintext
}
