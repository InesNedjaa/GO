package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()
	
		if err := conn.WriteMessage(websocket.TextMessage, []byte("OK")); err != nil {
			log.Println("Write error:", err)
			
		}
	}


func StartWSserver(metadata *OobMetadata) {
	http.HandleFunc("/ws", handleWebSocket)

	server := &http.Server{
		Addr:   metadata.Address + ":" + metadata.Port,
	}
	fmt.Printf("Starting server on port %s...\n",metadata.Address)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func main() {
	metadata := OobMetadata{
		Name:           "oob_module_2",
		Address:        "127.0.0.1",
		Port:           "8082",
		Connection_type: "WS",
	}

	err := saveMetadata(metadata)
	if err != nil {
		fmt.Println("Error saving metadata:", err)
	} else {
		fmt.Println("Module metadata saved successfully")
	}

	StartWSserver(&metadata)
}