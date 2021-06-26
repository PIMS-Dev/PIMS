package PIMS_utils

import (
	"bytes"
	"encoding/binary"
)

func ConvertBytesToUint32(data [4]byte) uint32 {
	var temp uint32
	bytesBuffer := bytes.NewBuffer(data[:4])
	binary.Read(bytesBuffer, binary.BigEndian, &temp)
	return uint32(temp)
}

func ConvertUint32ToBytes(data uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &data)
	return bytesBuffer.Bytes()[:4]
}

func ConvertBytesToUint64(data [8]byte) uint64 {
	var temp uint64
	bytesBuffer := bytes.NewBuffer(data[:8])
	binary.Read(bytesBuffer, binary.BigEndian, &temp)
	return uint64(temp)
}

func ConvertUint64ToBytes(data uint64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &data)
	return bytesBuffer.Bytes()[:8]
}

func ConvertBoolToByte(data bool) byte {
	var buffer byte
	if data {
		buffer = 1
	}
	return buffer
}