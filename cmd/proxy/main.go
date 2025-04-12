package main

import (
	"fmt"
	"log"
	"net"
	pb "go-proxy/api/proxy_service"
	"go-proxy/internal/proxy"

	"google.golang.org/grpc"
)

const (
	proxy_server_addr = "127.0.0.1"
	proxy_server_port = "8888"
)



func StartProxyServer () error{
	listener, err := net.Listen("tcp",  fmt.Sprintf(":%s", proxy_server_port)) 
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return err
	}

	grpc:= grpc.NewServer()
	pb.RegisterProxyServiceServer(grpc, &proxy.ProxyServer{})

	log.Printf("Server is running on port %s", proxy_server_port)
	if err := grpc.Serve(listener); err != nil {
		log.Printf("error while serving : %s", err)
		return err
	}
	return nil 
}

func main() {

err := StartProxyServer()
if err != nil {
	log.Printf("Failed to start proxy server")
	
}

}