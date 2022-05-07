package utils

import (
	"io/ioutil"
	"os"
)

func ReadStdin() ([]byte, error) {
	return ioutil.ReadAll(os.Stdin)
}
