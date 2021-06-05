package main

import (
	"./PIMS/PIMS_utils"
	"fmt"
)

func main() {
	bytes := []byte{0,0,1,0}
	newInt := PIMS_utils.ConvertBytesToUint32(bytes)
	fmt.Println(newInt)
	newBytes := PIMS_utils.ConvertUint32ToBytes(newInt)
	fmt.Println(newBytes)
}
