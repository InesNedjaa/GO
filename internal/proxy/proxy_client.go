package proxy

import (
	"context"
	"fmt"
	"log"
	"go-proxy/api/proxy_service"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcClient proxy_service.ProxyServiceClient
	grpcConn   *grpc.ClientConn
	once       sync.Once
)

func InitGrpcClient () error{
	var err error
	once.Do(func() {
		grpcConn, err = grpc.NewClient("localhost:8888" , grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection to grpc server failed")
		
		return 
	}
	log.Println("Connected to grpc server")
	grpcClient = proxy_service.NewProxyServiceClient(grpcConn)
	})
	return err 

}

func ForwardRequest (data map[string]string , module_name string) error {

	err := InitGrpcClient()
	if err != nil {
		return err
	}
	
	ctx , cancel:= context.WithTimeout(context.Background() , 5*time.Second)
	defer cancel()

	response , err := grpcClient.RecieveServiceRequest(
		ctx ,
		&proxy_service.Request{Data: data,ModuleName: module_name},
		)
	if err != nil {
		log.Fatalf("Error calling Proxy Service: %v", err)
		return err
	} else {
		log.Println("Server response : " , response.Message)
	}
	return nil

}


func SendMetadata(metadata *proxy_service.MetadataRequest ) error {
	err := InitGrpcClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
   response, err := grpcClient.Metadata(ctx , metadata)
	if err != nil {
		log.Fatalf("Failed to send metadata: %v", err)
		return err
	}
	fmt.Println("Server Response:", response.Message)
	return nil
}

