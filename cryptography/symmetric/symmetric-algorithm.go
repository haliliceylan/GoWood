package symmetric

type SymmetricAlgorithm interface {
	SetPublicKey([]byte)
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}
