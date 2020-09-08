package authorization

import (
	"crypto/rsa"
	"io/ioutil"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

var (
	once      sync.Once
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

// LoadFiles .
func LoadFiles(privateFile, publicFile string) error {
	var err error
	once.Do(func() {
		err = loadFiles(privateFile, publicFile)
	})
	return err
}

func loadFiles(privateFile, publicFile string) error {
	privateByte, err := ioutil.ReadFile(privateFile)
	if err != nil {
		return err
	}

	publicByte, err := ioutil.ReadFile(publicFile)
	if err != nil {
		return err
	}

	return parseRSA(privateByte, publicByte)
}

func parseRSA(privateByte, publicByte []byte) error {
	var err error
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateByte)
	if err != nil {
		return err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicByte)
	if err != nil {
		return err
	}
	return nil
}
