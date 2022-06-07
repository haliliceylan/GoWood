package hash

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"github.com/haliliceylan/gowood/utils"
)

type HashFlags struct {
	algorithm string        // algortihm
	filename  string        // file
	str       string        // string
	stdin     bool          // pipe mode
	flagSet   *flag.FlagSet // flagset
}

func MakeHashFlags(fg *flag.FlagSet) *HashFlags {
	return &HashFlags{
		flagSet: fg,
	}
}

func (hf *HashFlags) Prepare() {
	hf.flagSet.StringVar(&hf.algorithm, "algorithm", "md5", "hash algorithm")
	hf.flagSet.StringVar(&hf.filename, "filename", "", "input file name")
	hf.flagSet.StringVar(&hf.str, "string", "", "input string")
	hf.flagSet.BoolVar(&hf.stdin, "stdin", false, "should get input from stdin")
}

func (hf *HashFlags) Do() error {
	var currentAlgorithm HashAlgorithm

	switch hf.algorithm {
	case "md5":
		currentAlgorithm = &MD5{}
	case "sha256":
		currentAlgorithm = &SHA256{}
	case "sha512":
		currentAlgorithm = &SHA512{}
	default:
		log.Fatalf("No Algorithm Name %s", hf.algorithm)
	}

	var input []byte

	if hf.filename != "" {
		input, _ = utils.ReadFile(hf.filename)
	}

	if hf.str != "" {
		input = []byte(hf.str)
	}

	if hf.stdin {
		input, _ = utils.ReadStdin()
	}

	result, _ := currentAlgorithm.Hash(input)

	fmt.Printf("%s\n", hex.EncodeToString(result))
	return nil
}
