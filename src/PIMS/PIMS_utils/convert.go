package PIMS_utils

import (
	"bytes"
	"encoding/binary"
)

func ConvertBytesToUint32(data []byte) uint32 {
	var temp uint32
	bytesBuffer := bytes.NewBuffer(data)
	binary.Read(bytesBuffer, binary.BigEndian, &temp)
	return uint32(temp)
}

func ConvertUint32ToBytes(data uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &data)
	return bytesBuffer.Bytes()
}
