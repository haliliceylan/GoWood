package block

import (
	"crypto/cipher"
	"log"

	cryptoBlowfish "golang.org/x/crypto/blowfish"
)

type Blowfish struct {
	publicKey cipher.Block
}

func (bf *Blowfish) SetPublicKey(key []byte) {
	var err error
	bf.publicKey, err = cryptoBlowfish.NewCipher(key)
	if err != nil {
		log.Fatalf("Error during the Public Key Generating reason: %s", err.Error())
	}
}

func (bf *Blowfish) Encrypt(input []byte) []byte {
	if len(input) != 8 {
		log.Fatal("Blowfish is only supporting 8 bytes text input")
	}
	output := make([]byte, len(input))
	bf.publicKey.Encrypt(output, input)
	return output
}

func (bf *Blowfish) Decrypt(input []byte) []byte {
	output := make([]byte, len(input))
	bf.publicKey.Decrypt(output, input)
	return output
}
