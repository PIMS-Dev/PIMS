package main

import (
	"fmt"
	"./PIMS/PIMS_base"
)

func main() {
	sha256 := PIMS_base.DoubleSHA256([]byte("test string"))
	md5 := PIMS_base.DoubleMD5([]byte("test string"))
	fmt.Printf("%x\n%x\n", sha256, md5)
	orig := "test string"
	encrypted, _ := PIMS_base.EncryptAES256CBC([]byte(orig), sha256, md5)
	fmt.Println(orig)
	fmt.Printf("%x\n", encrypted)
	decrypted, err := PIMS_base.DecryptAES256CBC(encrypted, sha256, md5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", string(decrypted))
}
