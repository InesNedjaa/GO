package internal

import (
	"context"
	pb "oob-connector-proxy/v2/api/power_mgmt"
	"oob-connector-proxy/v2/internal/proxy"
)
type PowerMgmtServer struct {
	pb.UnimplementedPower_MgmtServer
	
	
}

func (s *PowerMgmtServer) PowerOn(context context.Context, request *pb.PowerOnDeviceRequest) (*pb.PowerDeviceResponse, error) {
    proxy_request := &proxy.ProxyType{
		Data: map[string]string{
			"mac_addr": request.MacAddr,
		},
		Oob_module: proxy.OobModuleType{
			Name:           "OOB_module_1", 
			Addr:           "127.0.0.1",
			Port:           "8081",
			Connection_type: "HTTP",
		},
	}
	proxy.RecieveServiceRequest(proxy_request)
    return nil, nil
}

func ( s *PowerMgmtServer) PowerOff (context context.Context , request *pb.PowerDeviceRequest)(*pb.PowerDeviceResponse , error) {
	
	proxy_request := &proxy.ProxyType{
		Data: map[string]string{
			"host": request.Host,
			"private_key_path": request.PrivateKeyPath,
			"username": request.Username,
		},
		Oob_module: proxy.OobModuleType{
			Name:           "OOB_module_1", 
			Addr:           "127.0.0.1",
			Port:           "8081",
			Connection_type: "HTTP",
		},
	}
	proxy.RecieveServiceRequest(proxy_request)
    return nil, nil
}

func ( s *PowerMgmtServer) Rebot (context context.Context , request *pb.PowerDeviceRequest)(*pb.PowerDeviceResponse , error) {
	
	proxy_request := &proxy.ProxyType{
		Data: map[string]string{
			"host": request.Host,
			"private_key_path": request.PrivateKeyPath,
			"username": request.Username,
		},
		Oob_module: proxy.OobModuleType{
			Name:           "OOB_module_1", 
			Addr:           "127.0.0.1",
			Port:           "8081",
			Connection_type: "HTTP",
		},
	}
	proxy.RecieveServiceRequest(proxy_request)
    return nil, nil
}