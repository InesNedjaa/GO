package internal

import (
	"context"
	"log"
	pb "go-proxy/api/power_mgmt"
	"go-proxy/internal/proxy"
)
type PowerMgmtServer struct {
	pb.UnimplementedPower_MgmtServer
	
	
}

func (s *PowerMgmtServer) PowerOn(context context.Context, request *pb.PowerOnDeviceRequest) (*pb.PowerDeviceResponse, error) {
    
	data:= map[string]string{"mac_addr": request.MacAddr,}
	oob_module_name:= "module_1"
	err := proxy.ForwardRequest(data,oob_module_name )
	if err != nil {
		log.Fatalf("Error forwarding request: %v", err)
		return nil ,err
	}
    return nil, nil
}

func ( s *PowerMgmtServer) PowerOff (context context.Context , request *pb.PowerDeviceRequest)(*pb.PowerDeviceResponse , error) {
	
	
	data:= map[string]string{
		"host": request.Host,
		"private_key_path": request.PrivateKeyPath,
		"username": request.Username,
	}
	oob_module_name:= "module_1"
	err := proxy.ForwardRequest(data,oob_module_name )
	if err != nil {
		log.Fatalf("Error forwarding request: %v", err)
		return nil ,err
	}   
	 return nil, nil
}

func ( s *PowerMgmtServer) Rebot (context context.Context , request *pb.PowerDeviceRequest)(*pb.PowerDeviceResponse , error) {
	
	data:= map[string]string{
		"host": request.Host,
		"private_key_path": request.PrivateKeyPath,
		"username": request.Username,
	}
	oob_module_name:= "module_1"
	err := proxy.ForwardRequest(data,oob_module_name )
	if err != nil {
		log.Fatalf("Error forwarding request: %v", err)
		return nil ,err
	}
		return nil, nil
}