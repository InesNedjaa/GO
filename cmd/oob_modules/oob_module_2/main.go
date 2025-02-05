package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"oob-connector-proxy/v2/api/proxy_service"
	"oob-connector-proxy/v2/internal/proxy"
	"os"

	"github.com/gorilla/websocket"
)

type OobMetadata struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	Port            string `json:"port"`
	Connection_type string `json:"connection_type"`
}

func GetMetadata() (*OobMetadata ,error ){
	file, err := os.Open("config/metadata_oob_m2.json")
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
	server_addr := fmt.Sprintf("%s:%s", metadata.Address, metadata.Port)
	http.HandleFunc("/ws", handleWebSocket)

	server := &http.Server{
		Addr:   server_addr,
	}
	fmt.Println("Server is listening on " , server_addr)
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
	StartWSserver(metadata)
	

}