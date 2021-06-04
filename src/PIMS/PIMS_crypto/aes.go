package PIMS_crypto

import (
	"errors"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func EncryptAES256CBC(data []byte, key [32]byte, iv [16]byte) ([]byte, error) {
	block, err := aes.NewCipher(key[:32])

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	correctData := bytes.Repeat([]byte{0}, blockSize)
	data = append(PKCS7Padding(data, blockSize), correctData...)
	blockMode := cipher.NewCBCEncrypter(block, iv[:16])

	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)

	return crypted, nil
}

func DecryptAES256CBC(data []byte, key [32]byte, iv [16]byte) ([]byte, error) {
	block, err := aes.NewCipher(key[:32])
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:16])

	decrypted := make([]byte, len(data))
	blockMode.CryptBlocks(decrypted, data)
	if !bytes.Equal(decrypted[(len(decrypted) - blockSize):], bytes.Repeat([]byte{0}, blockSize)) {
		return nil, errors.New("Decrypt key wrong.")
	}

	decrypted = PKCS7UnPadding(decrypted[:(len(decrypted) - blockSize)])
	return decrypted, nil
}

func PKCS7Padding(data []byte, blocksize int) []byte {
	padding := blocksize - len(data) % blocksize
	padbyte := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padbyte...)
}

func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
