package main

import (
	"os"
	"fmt"
	//"net"
	"io/ioutil"
	"encoding/json"
	//"PIMS/PIMS_crypto"
	//"PIMS/PIMS_utils"
)

var (
	config configFileStruct
)

func main() {
	fmt.Println("PIMS server starting...")
	err := createRootDirectory()
	err = readConfig()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Server exit.")
		os.Exit(1)
	}
}

type configFileStruct struct {
	ServerIp string `json:"server_ip"`
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