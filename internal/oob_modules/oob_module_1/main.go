package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OobMetadata struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	Port            string `json:"port"`
	Connection_type string `json:"connection_type"`
}

func saveMetadata(metadata OobMetadata) error {
	file, err := os.OpenFile("metadata.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
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

func startHTTPServer (metadata *OobMetadata) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK") 
	})
	server := &http.Server {
		Addr: metadata.Address + ":" + metadata.Port}
		fmt.Printf("Starting server on port %s...\n",metadata.Address)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

    
}

func main() {
	metadata:= OobMetadata{
		Name:           "oob_module_1",
		Address:        "127.0.0.1",
		Port:           "8081",
		Connection_type: "HTTP",
	}

	err := saveMetadata(metadata)
	if err != nil {
		fmt.Println("Error saving metadata:", err)
	} else {
		fmt.Println("Module metadata saved successfully")
	}
	startHTTPServer(&metadata)
}