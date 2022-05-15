package main

import (
	"flag"
	"os"

	"github.com/haliliceylan/gowood/cryptography/asymmetric"
	"github.com/haliliceylan/gowood/cryptography/hash"
	"github.com/haliliceylan/gowood/cryptography/symmetric"
)

func main() {
	hashFlagSet := flag.NewFlagSet("hash", flag.ExitOnError)
	hf := hash.MakeHashFlags(hashFlagSet)
	SymmetricFlagSet := flag.NewFlagSet("symmetric", flag.ExitOnError)
	sf := symmetric.MakeSymmetricFlags(SymmetricFlagSet)
	AsymmetricFlagSet := flag.NewFlagSet("asymmetric", flag.ExitOnError)
	af := asymmetric.MakeAsymmetricFlags(AsymmetricFlagSet)

	switch os.Args[1] {
	case "hash":
		hf.Prepare()
		hashFlagSet.Parse(os.Args[2:])
		hf.Do()
	case "symmetric":
		sf.Prepare()
		SymmetricFlagSet.Parse(os.Args[2:])
		sf.Do()
	case "asymmetric":
		af.Prepare()
		AsymmetricFlagSet.Parse(os.Args[2:])
		af.Do()
	}
}
