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

func (af *SymmetricFlags) Prepare() {
	af.flagSet.StringVar(&af.algorithm, "algorithm", "aes", "hash algorithm")
	af.flagSet.StringVar(&af.publicKey, "public", "", "public key should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256")
	af.flagSet.StringVar(&af.mode, "mode", "encrypt", "encrypt/decrypt")
	af.flagSet.StringVar(&af.filename, "filename", "", "input file name")
	af.flagSet.StringVar(&af.str, "string", "", "input string")
	af.flagSet.BoolVar(&af.stdin, "stdin", false, "should get input from stdin")
}

func (af *SymmetricFlags) Do() error {
	var currentAlgorithm SymmetricAlgorithm

	switch af.algorithm {
	case "aes":
		currentAlgorithm = &block.AES{}
	case "blowfish":
		currentAlgorithm = &block.Blowfish{}
	default:
		log.Fatalf("No Algorithm Name %s", af.algorithm)
	}

	if af.publicKey == "" {
		log.Fatalf("you must provide a public key")
	}

	currentAlgorithm.SetPublicKey([]byte(af.publicKey))

	var input []byte
	if af.filename != "" {
		input, _ = utils.ReadFile(af.filename)
	}

	if af.str != "" {
		input = []byte(af.str)
	}

	if af.stdin {
		input, _ = utils.ReadStdin()
	}

	var result []byte

	switch af.mode {
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
