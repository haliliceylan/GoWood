package hash

type HashAlgorithm interface {
	Hash([]byte) ([]byte, error)
}
