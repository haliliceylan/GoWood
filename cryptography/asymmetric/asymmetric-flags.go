package asymmetric

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/haliliceylan/gowood/utils"
)

type AsymmetricFlags struct {
	algorithm         string        // algortihm
	filename          string        // file
	str               string        // string
	mode              string        // publicKey
	clearStringVerify string        // clearStringVerify
	stdin             bool          // pipe mode
	createKey         bool          // create private public key
	flagSet           *flag.FlagSet // flagset
}

func MakeAsymmetricFlags(fg *flag.FlagSet) *AsymmetricFlags {
	return &AsymmetricFlags{
		flagSet: fg,
	}
}

func (af *AsymmetricFlags) Prepare() {
	af.flagSet.StringVar(&af.algorithm, "algorithm", "rsa", "hash algorithm")
	af.flagSet.BoolVar(&af.createKey, "createKey", false, "create/override new key value")
	af.flagSet.StringVar(&af.mode, "mode", "encrypt", "encrypt/decrypt/sign/verify")
	af.flagSet.StringVar(&af.filename, "filename", "", "input file name")
	af.flagSet.StringVar(&af.str, "string", "", "input string")
	af.flagSet.StringVar(&af.clearStringVerify, "clearStringVerify", "", "plainText for verify (required for verify)")
	af.flagSet.BoolVar(&af.stdin, "stdin", false, "should get input from stdin")
}

func (af *AsymmetricFlags) Do() error {
	var currentAlgorithm AsymmetricAlgorithm

	switch af.algorithm {
	case "rsa":
		currentAlgorithm = &RSA{}
	default:
		log.Fatalf("No Algorithm Name %s", af.algorithm)
	}

	if af.createKey {
		currentAlgorithm.CreateKeys()
		os.Exit(0)
	}

	currentAlgorithm.ReadKeys()

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
	case "sign":
		result = currentAlgorithm.Sign(input)
		fmt.Printf("%s", base64.URLEncoding.EncodeToString(result))
	case "verify":
		if af.clearStringVerify == "" {
			log.Fatal("clearStringVerify is required for verify operation")
		}
		input, _ = base64.URLEncoding.DecodeString(string(input))
		result := currentAlgorithm.Verify([]byte(af.clearStringVerify), input)
		var resultStr string
		if result {
			resultStr = "verified"
		} else {
			resultStr = "not verified"
		}
		fmt.Printf("%s\t%s", af.clearStringVerify, resultStr)

	}

	fmt.Println()

	return nil
}
