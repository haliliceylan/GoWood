package block

import (
	"strings"
	"testing"
)

func TestDES(t *testing.T) {
	algorithm := DES{}

	algorithm.SetPublicKey([]byte(strings.Repeat("s", 8)))

	plainText := "test"
	cipherText := string(algorithm.Encrypt([]byte(plainText)))

	plainTextResult := algorithm.Decrypt([]byte(cipherText))

	if plainText != string(plainTextResult) {
		t.Fatal("DES is not working properly")
	}
}
