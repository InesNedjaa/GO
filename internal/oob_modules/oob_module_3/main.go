package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

type OobMetadata struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	Port            string `json:"port"`
	Connection_type string `json:"connection_type"`
}

func saveMetadata(metadata OobMetadata) error {
	file, err := os.OpenFile("metadata.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(metadata)
	if err != nil {
		return err
	}

	return nil
}

func StartIPCserver (metadata OobMetadata) {
	server_addr := fmt.Sprintf("%s:%s", metadata.Address, metadata.Port)
	listener, err := net.Listen("tcp", server_addr)     //conn, err := net.Dial("unix", socket_path) for communication btw services in the same machine
	if err != nil {
		log.Fatal("Error creating Unix socket listener:", err)
	}

	fmt.Println("Server is listening on", server_addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			 
		}
		_, err = conn.Write([]byte("OK"))
	    if err != nil {
		log.Println("Error writing to connection:", err)
	}
		
	}
	

}

func main() {
	metadata := OobMetadata{
		Name:           "oob_module_3",
		Address:        "127.0.0.1",
		Port:           "8083",
		Connection_type: "IPC",
	}

	err := saveMetadata(metadata)
	if err != nil {
		fmt.Println("Error saving metadata:", err)
	} else {
		fmt.Println("Module metadata saved successfully")
	}
	StartIPCserver(metadata)
}