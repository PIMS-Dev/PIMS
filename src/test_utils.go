package main

import (
	"PIMS/PIMS_utils"
	"fmt"
)

func main() {
	bytes := [4]byte{0,0,1,0}
	newInt := PIMS_utils.ConvertBytesToUint32(bytes)
	fmt.Println(newInt)
	newBytes := PIMS_utils.ConvertUint32ToBytes(newInt)
	fmt.Println(newBytes)
	bytes2 := [8]byte{0,0,1,0,0,0,0,0}
	newInt2 := PIMS_utils.ConvertBytesToUint64(bytes2)
	fmt.Println(newInt2)
	newBytes2 := PIMS_utils.ConvertUint64ToBytes(newInt2)
	fmt.Println(newBytes2)
}
