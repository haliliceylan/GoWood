package hash

import (
	"encoding/hex"
	"testing"
)

func TestSha512(t *testing.T) {
	sha512 := SHA512{}
	expected := "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"
	result, _ := sha512.Hash([]byte("test"))
	resultStr := hex.EncodeToString(result)
	if resultStr != expected {
		t.Errorf("sha512 Function is not working sha512('test'), expected '%s', got '%s'", expected, resultStr)
	}
}
