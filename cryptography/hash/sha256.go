package hash

import cryptoSha256 "crypto/sha256"

type SHA256 struct{}

func (sha *SHA256) Hash(input []byte) ([]byte, error) {
	result := cryptoSha256.Sum256(input)
	return result[:], nil
}
