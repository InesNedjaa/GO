package internal

import (
	"context"
	pb "oob-connector-proxy/v2/api/monitoring_mgmt"
	"oob-connector-proxy/v2/internal/proxy"
	"strconv"

	"google.golang.org/grpc"
)
type MonitoringMgmtServer struct {
	pb.UnimplementedMonitoringMgmtServer
	
}


func (s *MonitoringMgmtServer) GetMetrics (context context.Context , request *pb.MetricsRequest) (*pb.MetricsResponse , error) {
	   
	proxy_request := &proxy.ProxyType{
		Data: map[string]string{
			"device_id": strconv.Itoa(int(request.DeviceId)),
			
		},
		Oob_module: proxy.OobModuleType{
			Name:           "OOB_module_2", 
			Addr:           "127.0.0.1",
			Port:           "8082",
			Connection_type: "WS",
		},
	}
	proxy.RecieveServiceRequest(proxy_request)
    return nil, nil
	
}

func (s *MonitoringMgmtServer) GetLogs ( request *pb.LogsRequest , stream grpc.ServerStreamingServer[pb.LogResponse])  error {
	
	
	proxy_request := &proxy.ProxyType{
		Data: map[string]string{
			"device_id": strconv.Itoa(int(request.DeviceId)),
			
		},
		Oob_module: proxy.OobModuleType{
			Name:           "OOB_module_2", 
			Addr:           "127.0.0.1",
			Port:           "8082",
			Connection_type: "WS",
		},
	}
	proxy.RecieveServiceRequest(proxy_request)
    return nil 

}

