package main

import (
	"os"
	"fmt"
	"net"
	"bytes"
	"strconv"
	"io/ioutil"
	"encoding/json"
	//"PIMS/PIMS_crypto"
	"PIMS/PIMS_utils"
)

var (
	config configFileStruct
	PIMSProtocolVersion uint8 = 1
)

func main() {
	fmt.Println("PIMS server starting...")
	err := createRootDirectory()
	err = readConfig()
	handelMainError(err)

	listener, err := net.Listen("tcp", config.ServerIP + ":" + strconv.Itoa(int(config.ServerPort)))
	handelMainError(err)
	defer listener.Close()

	fmt.Println("Server started at " + config.ServerIP + ":" + strconv.Itoa(int(config.ServerPort)))
	for {
		conn, err := listener.Accept()
		handelMainError(err)
		go handelConnection(conn)
	}
}

func handelConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 22)
	recvLen, err := conn.Read(buffer)
	if handelConnectionError(err) {
		return
	}

	if recvLen != 22 {
		return
	}

	if !bytes.Equal(buffer[:12], []byte("PIMS-Protocol")[:12]) {
		return
	}

	protocolVersion := uint8(buffer[12])
	if protocolVersion > PIMSProtocolVersion {
		fmt.Println("Your PIMS server is out of date. Please update.")
		return
	}

	packageType := buffer[13]
	packageLength := PIMS_utils.ConvertBytesToUint64(buffer[14:])
	fmt.Println(packageType, packageLength)
}

func handelConnectionError(err error) bool {
	if err != nil {
		fmt.Println("Connection close: "+err.Error())
		return true
	}
	return false
}

func handelMainError(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("Server exit.")
		os.Exit(1)
	}
}

type configFileStruct struct {
	ServerIP string `json:"server_ip"`
	ServerPort uint16 `json:"server_port"`
	EnableRegistration bool `json:"enable_registration"`
	EnableCreateChat bool `json:"enable_create_chat"`
}

func readConfig() error {
	if !pathExist("./config.json") {
		file, err := os.Create("./config.json");
		if err != nil {
			return err
		}
		defer file.Close()

		config = configFileStruct{"0.0.0.0", 17865, true, true}
		data, err := json.MarshalIndent(&config, "", "\t")
		if err != nil {
			return err
		}

		file.Write(data)
	} else {
		file, err := os.Open("./config.json")
		if err != nil {
			return err
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		defer file.Close()

		err = json.Unmarshal(data, &config)
		if err != nil {
			return err
		}
	}

	return nil
}

func createRootDirectory() error {
	if  !pathExist("./account") {
		err := os.MkdirAll("./account", os.ModePerm)
		if err != nil {
			return err
		}
	}

	if  !pathExist("./chat") {
		err := os.MkdirAll("./chat", os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func pathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
