package internal

import (
	"context"
	"log"
	pb "oob-connector-proxy/v2/api/monitoring_mgmt"
	"oob-connector-proxy/v2/internal/proxy"
	"strconv"

	"google.golang.org/grpc"
)
type MonitoringMgmtServer struct {
	pb.UnimplementedMonitoringMgmtServer
	
}




func (s *MonitoringMgmtServer) GetMetrics (context context.Context , request *pb.MetricsRequest) (*pb.MetricsResponse , error) {
	   
	data := map[string]string{"device_id": strconv.Itoa(int(request.DeviceId)),}
	oob_module_name:= "oob_module_2"
	err := proxy.ForwardRequest(data,oob_module_name )
	if err != nil {
		log.Fatalf("Error forwarding request: %v", err)
		return nil ,err
	}
    return nil, nil
	
}

func (s *MonitoringMgmtServer) GetLogs ( request *pb.LogsRequest , stream grpc.ServerStreamingServer[pb.LogResponse])  error {
	
	data:= map[string]string{"device_id": strconv.Itoa(int(request.DeviceId)),}
	oob_module_name:= "oob_module_2"
	err := proxy.ForwardRequest(data,oob_module_name )
	if err != nil {
		log.Fatalf("Error forwarding request: %v", err)
		return err
	}
    return nil

}

