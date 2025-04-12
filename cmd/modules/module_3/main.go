package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"go-proxy/api/proxy_service"
	"go-proxy/internal/proxy"
	"os"
)

type Metadata struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	Port            string `json:"port"`
	Connection_type string `json:"connection_type"`
}

func GetMetadata() (*Metadata ,error ){
	file, err := os.Open("config/metadata_module3.json")
	if err != nil {
		log.Printf("Error in opening file")
		return nil , err
	}
	defer file.Close()

	var metadata Metadata
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&metadata)
	if err != nil {
		log.Printf("Error in decoding %v" , err)
		return nil , err
	}

	return &metadata , nil
}

func StartIPCserver (metadata *Metadata) {
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
	metadata ,err := GetMetadata()
	if err != nil {
		log.Printf("Error in getting server metadata : %v" , err)
		return 
	}
	proxy.SendMetadata(&proxy_service.MetadataRequest{
		Name: metadata.Name,
		Addr: metadata.Address,
		Port: metadata.Port,
		ConnectionType: metadata.Connection_type,
	})
	StartIPCserver(metadata)
	
	
}