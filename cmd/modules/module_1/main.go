package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	file, err := os.Open("config/metadata_module1.json")
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



func startHTTPServer (metadata *Metadata) {
	server_addr := fmt.Sprintf("%s:%s", metadata.Address, metadata.Port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK") 
	})
	server := &http.Server {
		Addr: server_addr}
		fmt.Println("Starting server on port " , server_addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
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
     startHTTPServer(metadata)

}