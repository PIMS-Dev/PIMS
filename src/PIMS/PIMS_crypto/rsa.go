package PIMS_crypto

import (
	"bytes"
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

func ConvertRSAPrivateKeyToBytes(privateKey *rsa.PrivateKey) ([]byte, error) {
	x509PrivateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	return x509PrivateKeyBytes, nil
}

func ConvertRSAPublicKeyToBytes(publicKey *rsa.PublicKey) ([]byte, error) {
	x509PublicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	return x509PublicKeyBytes, err
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

func EncryptRSA(data []byte, publicKey *rsa.PublicKey, bitLength uint32) ([]byte, error) {
	partLen := int(bitLength / 8 - 11)
	chunks := split(data, partLen)
	buffer := bytes.NewBuffer([]byte{})

	for _, chunk := range chunks {
		encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(encrypted)
	}

	return buffer.Bytes(), nil
}

func DecryptRSA(data []byte, privateKey *rsa.PrivateKey, bitLength uint32) ([]byte, error) {
	partLen := int(bitLength / 8)
	chunks := split(data, partLen)
	buffer := bytes.NewBuffer([]byte{})

	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(decrypted)
	}

	return buffer.Bytes(), nil
}

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}