package main

import (
	"fmt"
	"./PIMS/PIMS_crypto"
)

func main() {
	sha256 := PIMS_crypto.DoubleSHA256([]byte("test string"))
	md5 := PIMS_crypto.DoubleMD5([]byte("test string"))
	fmt.Printf("%x\n%x\n", sha256, md5)
	orig := "01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678912345678"
	encrypted, _ := PIMS_crypto.EncryptAES256CBC([]byte(orig), sha256, md5)
	fmt.Println(orig)
	fmt.Printf("%x\n", encrypted)
	decrypted, err := PIMS_crypto.DecryptAES256CBC(encrypted, sha256, md5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(decrypted))
	prikey, pubkey, _ := PIMS_crypto.GenerateRSA2048Key(				)
	sign, _ := PIMS_crypto.MakeRSASign([]byte(orig), prikey)
	pass := PIMS_crypto.VerifyRSASign([]byte(orig), sign, pubkey)
	fmt.Println(pass)
}
