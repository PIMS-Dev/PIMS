package PIMS_crypto

import (
	"crypto/sha256"
	"crypto/md5"
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
