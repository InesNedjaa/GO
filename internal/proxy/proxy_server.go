package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

type ProxyType struct {
	Data map [string]string
	Oob_module OobModuleType 
}

type OobModuleType struct {
	Name string
	Addr string
	Port string
	Connection_type string

}

func HTTPserver (request *ProxyType) error { 
	server_url := fmt.Sprintf("http://%s:%s/", request.Oob_module.Addr, request.Oob_module.Port)

	jsonData, err := json.Marshal(request.Data)
	if err != nil {
		log.Fatalf("Error parsing data: %v", err)
		return err
	}
	resp, err := http.Post(server_url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error sending POST request: %v", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("POST request sent successfully")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return err
	}
	fmt.Println("Response from HTTP Server :")
	fmt.Println(string(body))
   return nil
}


func WSserver (request *ProxyType) error{

	server_url := fmt.Sprintf("ws://%s:%s/ws", request.Oob_module.Addr, request.Oob_module.Port)	
	conn, _, err := websocket.DefaultDialer.Dial(server_url, nil)
	if err != nil {
		log.Fatalf("Error connecting to WebSocket server: %v", err)
		return nil
	} 
	defer conn.Close()
    fmt.Println("Connected to WebSocket server")
	
			_, message, err := conn.ReadMessage()
			if err != nil {
			
				log.Printf("Error reading message: %v", err)
				return err
			}
			fmt.Println("Response from WS Server :")
	        fmt.Println(string(message))
		

	return nil
}
func IPCserver (request *ProxyType) error{ 
	server_url := fmt.Sprintf("%s:%s", request.Oob_module.Addr, request.Oob_module.Port)
	conn, err := net.Dial("tcp", server_url)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
		return nil 
	}
    defer conn.Close()
	fmt.Println("Connected to IPC server")
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal("Error reading response:", err) 
	}
	fmt.Println("Response from IPC Server :")
	fmt.Println(string(buffer[:n]))
	return nil
}

func RecieveServiceRequest (request *ProxyType) error { 
     switch request.Oob_module.Connection_type {
	 case "HTTP" :
		HTTPserver(request)
	 case "WS" :
		WSserver(request)
	 case "IPC" :
		IPCserver(request)
	 }
	return nil 
} 