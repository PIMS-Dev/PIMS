package PIMS_crypto

import (
	"crypto"
	"crypto/sha512"
	"crypto/rsa"
	"crypto/x509"
	"crypto/rand"
)

func GenerateRSA2048Key() (*rsa.PrivateKey, *rsa.PublicKey, error){
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	return privateKey, &privateKey.PublicKey, err
}

func ConvertRSAPrivateKeyToBytes(privateKey *rsa.PrivateKey) []byte {
	x509PrivateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	return x509PrivateKeyBytes
}

func ConvertRSAPublicKeyToBytes(publicKey *rsa.PublicKey) []byte {
	x509PublicKeyBytes, _ := x509.MarshalPKIXPublicKey(publicKey)
	return x509PublicKeyBytes
}

func LoadRSAPrivateKey(keyData []byte) (*rsa.PrivateKey, error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(keyData)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func LoadRSAPublicKey(keyData []byte) (*rsa.PublicKey, error) {
	publicKeyInterface, err := x509.ParsePKIXPublicKey(keyData)
	if err != nil {
		return nil, err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey, err
}

func MakeRSASign(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha512.New()
	hash.Write(data)
	sum := hash.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, sum)
	if err != nil {
		return nil, err
	}
	return sign, nil
}

func VerifyRSASign(data []byte, sign []byte, publicKey *rsa.PublicKey) bool {
	hash := sha512.New()
	hash.Write(data)
	sum := hash.Sum(nil)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, sum, sign)
	return err == nil
}
