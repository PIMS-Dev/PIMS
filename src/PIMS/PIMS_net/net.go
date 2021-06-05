package PIMS_net

import (
	"net"
	"binary"
	"bytes"
	"errors"
)

func Receive(conn net.Conn) ([]byte, err) {
	lengthByte := make([]byte, 4)
	recvLen, err := conn.Read(lengthByte)
	if err != nil {
		return nil, err
	} else if recvLen != 4 {
		return nil, errors.New("Protocol error.")
	}

	var length uint32
	lengthByteBuffer := bytes.NewBuffer(lengthByte)
	err = binary.Read(lengthByteBuffer, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}

	data := make([]byte, length)
	buffer := make([]byte, 8192)
	recvLen, err = conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	data = append(data, buffer)

	for recbLen == 8192 {
		recvLen, err = conn.Read(buffer)
		if err != nil {
			return nil, err
		}
		data = append(data, buffer)
	}

	if recvLen != length {
		return nil, errors.New("Protocol error.")
	}
	return data, nil
}
