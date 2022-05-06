package hash

import (
	cryptoMd5 "crypto/md5"
)

type MD5 struct{}

func (md5 *MD5) Hash(input []byte) ([]byte, error) {
	result := cryptoMd5.Sum(input)
	return result[:], nil
}
