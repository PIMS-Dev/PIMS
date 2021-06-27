package PIMS_client

import (
	"net"
	"bytes"
	"strconv"
	"PIMS/PIMS_utils"
)

var (
	conn net.Conn = nil
	PIMSProtocolVersion uint8 = 1
)

func ConnectServer(serverIP string, serverPort uint16) error {
	conn, err := net.Dial("tcp", serverIP + ":" + strconv.Itoa(int(serverPort)))
	return err
}

func login(username string, password string) {
	
}

func buildPackageHeader(code byte, bodyLength uint64) [22]byte {
	var buffer bytes.Buffer
	buffer.Write([]byte("PIMS-Protocol")[:12])
	buffer.Write([]byte{PIMSProtocolVersion, code})
	buffer.Write(PIMS_utils.ConvertUint64ToBytes(bodyLength))
	return buffer.Bytes()[:22]
}