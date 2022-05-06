package main

import (
	"flag"
	"os"

	"github.com/haliliceylan/gowood/cryptography/hash"
)

func main() {
	hashFlagSet := flag.NewFlagSet("hash", flag.ExitOnError)
	hf := hash.MakeHashFlags(hashFlagSet)

	switch os.Args[1] {
	case "hash":
		hf.Prepare()
		hashFlagSet.Parse(os.Args[2:])
		hf.Do()
	}
}
