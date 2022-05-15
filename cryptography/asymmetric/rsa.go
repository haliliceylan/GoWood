package asymmetric

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	CryptoRSA "crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/haliliceylan/gowood/utils"
)

const RSA_PUBLIC_KEY_NAME = "rsa_public.key"
const RSA_PRIVATE_KEY_NAME = "rsa_private.key"

type RSA struct {
	privateKey *CryptoRSA.PrivateKey
}

func (r *RSA) CreateKeys() {

	reader := rand.Reader
	bitSize := 2048

	privateKeyFile, err := os.Create(RSA_PRIVATE_KEY_NAME)

	if err != nil {
		log.Fatal(err)
	}

	defer privateKeyFile.Close()
	publicKeyFile, err := os.Create(RSA_PUBLIC_KEY_NAME)

	if err != nil {
		log.Fatal(err)
	}

	defer publicKeyFile.Close()

	privateKey, err := CryptoRSA.GenerateKey(reader, bitSize)

	if err != nil {
		log.Fatal(err)
	}

	var privateKeyPemBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = pem.Encode(privateKeyFile, privateKeyPemBlock)

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.PublicKey

	var publicKeyPemBlock = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&publicKey),
	}

	err = pem.Encode(publicKeyFile, publicKeyPemBlock)

	if err != nil {
		log.Fatal(err)
	}
}

func (r *RSA) ReadKeys() {
	privateKeyFile, err := utils.ReadFile(RSA_PRIVATE_KEY_NAME)

	if err != nil {
		log.Fatal(err)
	}

	_, err = utils.ReadFile(RSA_PUBLIC_KEY_NAME)

	if err != nil {
		log.Fatal(err)
	}

	privatePemBlock, _ := pem.Decode(privateKeyFile)

	privPemBytes := privatePemBlock.Bytes

	parsedPrivateKey, err := x509.ParsePKCS1PrivateKey(privPemBytes)

	if err != nil {
		log.Fatal(err)
	}

	r.privateKey = parsedPrivateKey
}

func (r *RSA) Encrypt(input []byte) []byte {

	encryptedBytes, err := CryptoRSA.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&r.privateKey.PublicKey,
		input,
		nil)

	if err != nil {
		log.Fatal(err)
	}

	return encryptedBytes
}

func (r *RSA) Decrypt(input []byte) []byte {
	decryptedBytes, err := r.privateKey.Decrypt(nil, input, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		log.Fatal(err)
	}
	return decryptedBytes
}

func (r *RSA) Sign(input []byte) []byte {
	// Before signing, we need to hash our message
	// The hash is what we actually sign
	msgHash := sha256.New()
	_, err := msgHash.Write(input)
	if err != nil {
		log.Fatal(err)
	}
	msgHashSum := msgHash.Sum(nil)

	// In order to generate the signature, we provide a random number generator,
	// our private key, the hashing algorithm that we used, and the hash sum
	// of our message
	signature, err := rsa.SignPSS(rand.Reader, r.privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	return signature
}

func (r *RSA) Verify(input []byte, signature []byte) bool {

	// Before signing, we need to hash our message
	// The hash is what we actually sign
	msgHash := sha256.New()
	_, err := msgHash.Write(input)
	if err != nil {
		log.Fatal(err)
	}
	msgHashSum := msgHash.Sum(nil)
	// To verify the signature, we provide the public key, the hashing algorithm
	// the hash sum of our message and the signature we generated previously
	// there is an optional "options" parameter which can omit for now
	err = rsa.VerifyPSS(&r.privateKey.PublicKey, crypto.SHA256, msgHashSum, signature, nil)
	return err == nil
}
