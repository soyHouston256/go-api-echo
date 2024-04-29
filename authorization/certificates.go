package authorization

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"sync"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

func LoadFiles(privateFile string, publicFile string) error {
	var err error
	once.Do(func() {
		err = loadKeys(privateFile, publicFile)
	})
	return err
}

func loadKeys(privateFile string, publicFile string) error {
	privateBytes, err := os.ReadFile(privateFile)
	if err != nil {
		return err
	}
	publicBytes, err := os.ReadFile(publicFile)
	if err != nil {
		return err
	}
	return parseRSA(privateBytes, publicBytes)
}

func parseRSA(privateBytes, publicBytes []byte) error {
	var err error
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}
	return nil
}
