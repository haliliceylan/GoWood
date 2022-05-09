package hash

import (
	"encoding/hex"
	"testing"
)

func TestMD5Hash(t *testing.T) {
	md5 := MD5{}
	expected := "098f6bcd4621d373cade4e832627b4f6"
	result, _ := md5.Hash([]byte("test"))
	resultStr := hex.EncodeToString(result)
	if resultStr != expected {
		t.Errorf("MD5 Function is not working md5('test'), expected '%s', got '%s'", expected, resultStr)
	}
}
