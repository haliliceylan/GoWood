package asymmetric

type AsymmetricAlgorithm interface {
	SetPublicKey([]byte)
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}
