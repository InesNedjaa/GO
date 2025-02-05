package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"oob-connector-proxy/v2/api/proxy_service"
	"oob-connector-proxy/v2/internal/proxy"
	"os"
)

type OobMetadata struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	Port            string `json:"port"`
	Connection_type string `json:"connection_type"`
}

func GetMetadata() (*OobMetadata ,error ){
	file, err := os.Open("config/metadata_oob_m1.json")
	if err != nil {
		log.Printf("Error in opening file")
		return nil , err
	}
	defer file.Close()

	var metadata OobMetadata
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&metadata)
	if err != nil {
		log.Printf("Error in decoding %v" , err)
		return nil , err
	}

	return &metadata , nil
}



func startHTTPServer (metadata *OobMetadata) {
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
	proxy.SendMetadata(&proxy_service.Metadata{
		Name: metadata.Name,
		Addr: metadata.Address,
		Port: metadata.Port,
		ConnectionType: metadata.Connection_type,
	})
     startHTTPServer(metadata)

}