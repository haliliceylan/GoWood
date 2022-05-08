package main

import (
	"flag"
	"os"

	"github.com/haliliceylan/gowood/cryptography/hash"
	"github.com/haliliceylan/gowood/cryptography/symmetric"
)

func main() {
	hashFlagSet := flag.NewFlagSet("hash", flag.ExitOnError)
	hf := hash.MakeHashFlags(hashFlagSet)
	SymmetricFlagSet := flag.NewFlagSet("symmetric", flag.ExitOnError)
	af := symmetric.MakeSymmetricFlags(SymmetricFlagSet)

	switch os.Args[1] {
	case "hash":
		hf.Prepare()
		hashFlagSet.Parse(os.Args[2:])
		hf.Do()
	case "symmetric":
		af.Prepare()
		SymmetricFlagSet.Parse(os.Args[2:])
		af.Do()
	}
}
