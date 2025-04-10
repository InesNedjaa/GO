package main

import (
	"context"
	"log"
	"net/http"
	"grpc-gateway/api/service1"
	"grpc-gateway/api/service2"
	"grpc-gateway/api/service3"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
const (
	gateway_port = "8080"
	server_addr = "localhost:8888"
)

func main () {
	conn , err := grpc.NewClient(server_addr,  grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("connection failed")

	}
	defer conn.Close()
	mux := runtime.NewServeMux()

	err = service2.RegisterService2Handler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalf("Failed to register Power_Mgmt gateway: %v", err)
	}

	err = service3.RegisterService3Handler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalf("Failed to register Script_Mgmt gateway: %v", err)
	}

	err = service1.RegisterService1Handler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalf("Failed to register MonitoringMgmt gateway: %v", err)
	}
	gwServer := &http.Server{
		Addr:   ":" + gateway_port,
		Handler: mux,
	}    
	log.Printf("Starting GRPC Gateway server on port %s", gateway_port)
	log.Fatalln(gwServer.ListenAndServe())
} 