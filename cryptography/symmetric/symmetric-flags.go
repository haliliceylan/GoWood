package symmetric

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"

	"github.com/haliliceylan/gowood/cryptography/symmetric/block"
	"github.com/haliliceylan/gowood/utils"
)

type SymmetricFlags struct {
	algorithm string        // algortihm
	filename  string        // file
	str       string        // string
	publicKey string        // publicKey
	mode      string        // publicKey
	stdin     bool          // pipe mode
	flagSet   *flag.FlagSet // flagset
}

func MakeSymmetricFlags(fg *flag.FlagSet) *SymmetricFlags {
	return &SymmetricFlags{
		flagSet: fg,
	}
}

func (sf *SymmetricFlags) Prepare() {
	sf.flagSet.StringVar(&sf.algorithm, "algorithm", "aes", "hash algorithm")
	sf.flagSet.StringVar(&sf.publicKey, "public", "", "public key")
	sf.flagSet.StringVar(&sf.mode, "mode", "encrypt", "encrypt/decrypt")
	sf.flagSet.StringVar(&sf.filename, "filename", "", "input file name")
	sf.flagSet.StringVar(&sf.str, "string", "", "input string")
	sf.flagSet.BoolVar(&sf.stdin, "stdin", false, "should get input from stdin")
}

func (sf *SymmetricFlags) Do() error {
	var currentAlgorithm SymmetricAlgorithm

	switch sf.algorithm {
	case "aes":
		currentAlgorithm = &block.AES{}
	case "blowfish":
		currentAlgorithm = &block.Blowfish{}
	case "des":
		currentAlgorithm = &block.DES{}
	default:
		log.Fatalf("No Algorithm Name %s", sf.algorithm)
	}

	if sf.publicKey == "" {
		log.Fatalf("you must provide a public key")
	}

	currentAlgorithm.SetPublicKey([]byte(sf.publicKey))

	var input []byte
	if sf.filename != "" {
		input, _ = utils.ReadFile(sf.filename)
	}

	if sf.str != "" {
		input = []byte(sf.str)
	}

	if sf.stdin {
		input, _ = utils.ReadStdin()
	}

	var result []byte

	switch sf.mode {
	case "encrypt":
		result = currentAlgorithm.Encrypt(input)
		fmt.Printf("%s", base64.URLEncoding.EncodeToString(result))
	case "decrypt":
		input, _ = base64.URLEncoding.DecodeString(string(input))
		result = currentAlgorithm.Decrypt(input)
		fmt.Printf("%s", string(result))
	}

	fmt.Println()

	return nil
}
