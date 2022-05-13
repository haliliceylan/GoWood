package block

import (
	"strings"
	"testing"
)

func TestBlowFish(t *testing.T) {
	algorithm := Blowfish{}

	algorithm.SetPublicKey([]byte(strings.Repeat("s", 32)))

	plainText := "testtest"
	cipherText := string(algorithm.Encrypt([]byte(plainText)))

	plainTextResult := algorithm.Decrypt([]byte(cipherText))

	if plainText != string(plainTextResult) {
		t.Fatal("BlowFish is not working properly")
	}
}
