package hash

import cryptoSha512 "crypto/sha512"

type SHA512 struct{}

func (sha *SHA512) Hash(input []byte) ([]byte, error) {
	result := cryptoSha512.Sum512(input)
	return result[:], nil
}
