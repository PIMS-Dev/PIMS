package PIMS_base

import (
	"errors"
	"bytes"
	"crypto/sha256"
	"crypto/md5"
	"crypto/aes"
	"crypto/cipher"
)

func DoubleSHA256(data []byte) [32]byte {
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:32])
	return hash2
}

func DoubleMD5(data []byte) [16]byte {
	hash1 := md5.Sum(data)
	hash2 := md5.Sum(hash1[:16])
	return hash2
}

func EncryptAES256CBC(data []byte, key [32]byte, iv [16]byte) ([]byte, error) {
	key_slice := make([]byte, 32)
	iv_slice := make([]byte, 16)
	copy(key_slice, key[:32])
	copy(iv_slice, iv[:16])

	block, err := aes.NewCipher(key_slice)

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	correctData := bytes.Repeat([]byte{0}, blockSize)
	data = append(PKCS7Padding(data, blockSize), correctData...)
	//data = PKCS7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv_slice)

	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)

	return crypted, nil
}

func DecryptAES256CBC(data []byte, key [32]byte, iv [16]byte) ([]byte, error) {
	key_slice := make([]byte, 32)
	iv_slice := make([]byte, 16)
	copy(key_slice, key[:32])
	copy(iv_slice, iv[:16])

	block, err := aes.NewCipher(key_slice)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv_slice)

	decrypted := make([]byte, len(data))
	blockMode.CryptBlocks(decrypted, data)
	if !bytes.Equal(decrypted[(len(decrypted) - blockSize):], bytes.Repeat([]byte{0}, blockSize)) {
		err = errors.New("Decrypt key wrong.")
		return nil, err
	}

	decrypted = PKCS7UnPadding(decrypted[:(len(decrypted) - blockSize)])
	//decrypted = PKCS7UnPadding(decrypted)
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
