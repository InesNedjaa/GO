package proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"

	pb "oob-connector-proxy/v2/api/proxy_service"

	"github.com/gorilla/websocket"
)

type ProxyType struct {
	Data map [string]string
	Oob_module_name string
}

type OobMetaData struct {
	Name string
	Addr string
	Port string
	Connection_type string

}
type ProxyServer struct {
	pb.UnimplementedProxyServiceServer
	Servers []OobMetaData 
	serversMutex sync.RWMutex
}


func (s *ProxyServer)OobMetadata(ctx context.Context, req *pb.Metadata) (*pb.Response, error){
	server_metadat:= OobMetaData{
		Name: req.Name,
		Addr: req.Addr,
		Port: req.Port,
		Connection_type: req.ConnectionType }
	s.Servers = append(s.Servers, server_metadat)
	return &pb.Response{Message: "Metadata recieved successfully"} , nil

}
func findServerByName(server_name string , Servers []OobMetaData ) *OobMetaData {
	
	//serversMutex.RLock() 
	//defer serversMutex.RUnlock()
	
	for i := range Servers {
		if Servers[i].Name == server_name {
			return &Servers[i]
		}
	}
	return nil
}
func HTTPserver (oob_module *OobMetaData , data map [string]string ) (string ,error ){ 
	var  response string 
	server_url := fmt.Sprintf("http://%s:%s/", oob_module.Addr, oob_module.Port)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error parsing data: %v", err)
		return response ,err
	}
	resp, err := http.Post(server_url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error sending POST request: %v", err)
		return response ,err
	}
	defer resp.Body.Close()
	fmt.Println("POST request sent successfully")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return response ,err
	}
	response = string(body)
	fmt.Println("Response from HTTP Server :")
	fmt.Println(response)
   return response ,nil
}


func WSserver (oob_module *OobMetaData ) (string ,error){
   var response string 
	server_url := fmt.Sprintf("ws://%s:%s/ws", oob_module.Addr, oob_module.Port)	
	conn, _, err := websocket.DefaultDialer.Dial(server_url, nil)
	if err != nil {
		log.Fatalf("Error connecting to WebSocket server: %v", err)
		return response, err
	} 
	defer conn.Close()
    fmt.Println("Connected to WebSocket server")
	
	_, message, err := conn.ReadMessage()
	if err != nil {
	
		log.Printf("Error reading message: %v", err)
		return response ,err
	}
	response = string(message)
	fmt.Println("Response from WS Server :")
	fmt.Println(response)
		

	return response,nil
}
func IPCserver (oob_module *OobMetaData  ) (string ,error){ 
	var response string 
	server_url := fmt.Sprintf("%s:%s", oob_module.Addr, oob_module.Port)
	conn, err := net.Dial("tcp", server_url)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
		return response ,err
	}
    defer conn.Close()
	fmt.Println("Connected to IPC server")
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal("Error reading response:", err) 
		return response , err
	}
	response = string(buffer[:n])
	fmt.Println("Response from IPC Server :")
	fmt.Println(response)
	return response ,nil
}


func (s *ProxyServer)RecieveServiceRequest (ctx context.Context ,request *pb.Request) (*pb.Response, error) { 

	oob_module := findServerByName(request.OobModuleName , s.Servers)
	if oob_module == nil {
		log.Printf("server not found")
		return nil ,errors.New("server not found")
	}
     switch oob_module.Connection_type {
	 case "HTTP" :
		
			response , err:=HTTPserver(oob_module , request.Data)
			if err != nil {
				log.Println("Error in connecting to HTTP server")
				return nil,err
			}
			return &pb.Response{Message: response} , nil
	
		
	 case "WS" :
			response ,err := WSserver(oob_module)
			if err != nil {
				log.Println("Error in connecting to WS server")
				return nil,err
			}
			return &pb.Response{Message: response} , nil
		
		
	 case "IPC" :
			response , err := IPCserver(oob_module)
			if err != nil {
				log.Println("Error in connecting to HTTP server")
				return nil,err
			}
			return &pb.Response{Message: response} , nil
		
	 }
	return &pb.Response{Message: "Server not found"} , nil 
} 

