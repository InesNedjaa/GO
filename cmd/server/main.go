package main

import (
	"fmt"
	"log"
	"net"
	"grpc-gateway/api/service1"
	"grpc-gateway/api/service2"
	"grpc-gateway/api/service3"
	"grpc-gateway/internal"

	"google.golang.org/grpc"
)
const (
	DefaultPort = "8888"
)

 

func StartServer () error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", DefaultPort))
	if err != nil {
		log.Printf("failed to listen on port %s err : %s", DefaultPort, err)
		return err
	}
	grpc := grpc.NewServer()
	service2.RegisterService2Server(grpc, &internal.Service2Server{})
	service3.RegisterService3Server(grpc, &internal.Service3Server{})
	service1.RegisterService1Server(grpc, &internal.Service1Server{})
	log.Printf("Server is running on port %s", DefaultPort)
	if err := grpc.Serve(listener); err != nil {
		log.Printf("error while serving : %s", err)
		return err
	}
	return nil

}

func main () {
	
	if err := StartServer(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} 
}