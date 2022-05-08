package block

import (
	"bytes"
	"crypto/cipher"
	"errors"
	"log"

	cryptoDES "crypto/des"
)

type DES struct {
	publicKey cipher.Block
}

func (des *DES) SetPublicKey(key []byte) {
	var err error
	des.publicKey, err = cryptoDES.NewCipher(key)
	if err != nil {
		log.Fatalf("Error during the Public Key Generating reason: %s", err.Error())
	}
}

func (des DES) Pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (des DES) PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func (des *DES) Encrypt(input []byte) []byte {
	bs := des.publicKey.BlockSize()
	input = des.Pkcs5Padding(input, bs)
	if len(input)%bs != 0 {
		log.Fatalf("%s", errors.New("need a multiple of the blocksize"))
	}

	out := make([]byte, len(input))
	dst := out
	for len(input) > 0 {
		//Encrypt encrypts the first block and saves the result to dst
		des.publicKey.Encrypt(dst, input[:bs])
		input = input[bs:]
		dst = dst[bs:]
	}
	return out
}

func (des *DES) Decrypt(input []byte) []byte {
	out := make([]byte, len(input))
	dst := out
	bs := des.publicKey.BlockSize()
	if len(input)%bs != 0 {
		log.Fatalf("%s", errors.New("crypto/cipher: input not full blocks"))
	}
	for len(input) > 0 {
		des.publicKey.Decrypt(dst, input[:bs])
		input = input[bs:]
		dst = dst[bs:]
	}
	out = des.PKCS5UnPadding(out)
	return out
}
