package block

import (
	"strings"
	"testing"
)

func TestAES256with32BytesKey(t *testing.T) {
	algorithm := AES{}

	algorithm.SetPublicKey([]byte(strings.Repeat("s", 32)))

	plainText := "test"
	cipherText := string(algorithm.Encrypt([]byte(plainText)))

	plainTextResult := algorithm.Decrypt([]byte(cipherText))

	if plainText != string(plainTextResult) {
		t.Fatal("AES256 is not working properly")
	}
}

func TestAES192with24BytesKey(t *testing.T) {
	algorithm := AES{}

	algorithm.SetPublicKey([]byte(strings.Repeat("s", 24)))

	plainText := "test"
	cipherText := string(algorithm.Encrypt([]byte(plainText)))

	plainTextResult := algorithm.Decrypt([]byte(cipherText))

	if plainText != string(plainTextResult) {
		t.Fatal("AES192 is not working properly")
	}
}

func TestAES128with16BytesKey(t *testing.T) {
	algorithm := AES{}

	algorithm.SetPublicKey([]byte(strings.Repeat("s", 16)))

	plainText := "test"
	cipherText := string(algorithm.Encrypt([]byte(plainText)))

	plainTextResult := algorithm.Decrypt([]byte(cipherText))

	if plainText != string(plainTextResult) {
		t.Fatal("AES128 is not working properly")
	}
}
