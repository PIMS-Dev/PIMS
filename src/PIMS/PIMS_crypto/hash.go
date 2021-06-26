package PIMS_crypto

import (
	"bytes"
	"crypto/sha256"
	"crypto/md5"
)

func SHA256(data []byte) [32]byte {
	hash := sha256.Sum256(data)
	return hash
}

func SHA256WithSalt(data []byte, salt []byte) [32]byte {
	hash := sha256.Sum256(bytes.Join([][]byte{data, salt}, []byte{}))
	return hash
}

func MD5(data []byte) [16]byte {
	hash := md5.Sum(data)
	return hash
}
