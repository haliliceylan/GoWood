package main

import (
	"flag"
	"os"

	"github.com/haliliceylan/gowood/cryptography/asymmetric"
	"github.com/haliliceylan/gowood/cryptography/hash"
)

func main() {
	hashFlagSet := flag.NewFlagSet("hash", flag.ExitOnError)
	hf := hash.MakeHashFlags(hashFlagSet)
	AsymmetricFlagSet := flag.NewFlagSet("asymmetric", flag.ExitOnError)
	af := asymmetric.MakeAsymmetricFlags(AsymmetricFlagSet)

	switch os.Args[1] {
	case "hash":
		hf.Prepare()
		hashFlagSet.Parse(os.Args[2:])
		hf.Do()
	case "asymmetric":
		af.Prepare()
		AsymmetricFlagSet.Parse(os.Args[2:])
		af.Do()
	}
}
