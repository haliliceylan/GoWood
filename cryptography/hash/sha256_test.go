package hash

import (
	"encoding/hex"
	"testing"
)

func TestSha256(t *testing.T) {
	sha256 := SHA256{}
	expected := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	result, _ := sha256.Hash([]byte("test"))
	resultStr := hex.EncodeToString(result)
	if resultStr != expected {
		t.Errorf("sha256 Function is not working sha256('test'), expected '%s', got '%s'", expected, resultStr)
	}
}
