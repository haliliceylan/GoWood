package asymmetric

type AsymmetricAlgorithm interface {
	CreateKeys()
	ReadKeys()
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
	Sign([]byte) []byte
	Verify([]byte, []byte) bool
}
