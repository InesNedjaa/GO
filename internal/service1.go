package internal

import (
	"context"
	"fmt"
	"io"
	"math"
	pb "grpc-gateway/api/service1"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"google.golang.org/grpc"
)
type Service1Server struct {
	pb.UnimplementedService1Server
	
}


func (s *Service1Server) GetMetrics (context context.Context , request *pb.MetricsRequest) (*pb.MetricsResponse , error) {
	    var cpu_usage_percentage float64
		var memory_usage_percentage float64
		var network_bandwidth uint64
		var errorMessages []string
		cpu_usage , err := cpu.Percent(5*time.Second , false)
		if err != nil { 
			errorMessages = append(errorMessages, fmt.Sprintf("Error in getting CPU usage: %v", err))
			
		} else {cpu_usage_percentage = math.Round(cpu_usage[0]*100) / 100 }
	
		memory_usage , err := mem.VirtualMemory()
		if err != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("Error in getting memory usage: %v", err))
		} else {
		memory_usage_percentage = float64(memory_usage.Used) / float64(memory_usage.Total) * 100 
		memory_usage_percentage = math.Round(memory_usage_percentage*100) / 100 }
	
	
		initialStats, err := net.IOCounters(false)
		if err != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("Error fetching initial network stats: %v", err))
			
		} 
		time.Sleep(1 * time.Second)
		finalStats, err := net.IOCounters(false) 
		if err != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("Error fetching final network stats: %v", err))
			
		}
		if len(initialStats) == 0 || len(finalStats) == 0 {
			errorMessages = append(errorMessages, "No network interfaces found.")
			
		} else {
		bytesSent := finalStats[0].BytesSent - initialStats[0].BytesSent
		bytesRecv := finalStats[0].BytesRecv - initialStats[0].BytesRecv
		network_bandwidth = bytesSent + bytesRecv }
		var finalError error
	    if len(errorMessages) > 0 {
		finalError = fmt.Errorf("Metrics retrieval encountered errors: %v", errorMessages)
	}
 	return &pb.MetricsResponse{CpuUsage: float32(cpu_usage_percentage) , MemoryUsage: float32(memory_usage_percentage), NetworkBandwidth: network_bandwidth} , finalError
	
}

func (s *Service1Server) GetLogs ( request *pb.LogsRequest , stream grpc.ServerStreamingServer[pb.LogResponse])  error {
	var file_path string 
    switch request.GetLogType(){
	case pb.LogType_System_logs : file_path="/var/log/syslog"
	case pb.LogType_Kernel_logs : file_path ="/var/log/kern.log"
	case pb.LogType_Auth_logs : file_path ="/var/log/auth.log"
	}
	file, err := os.OpenFile(file_path, os.O_RDONLY , 0)
	if err != nil { 
		return err
	}
	defer file.Close() 
	buffer := make([]byte, 1024) 
    for {
        content, err := file.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break 
            }
            return fmt.Errorf("error reading file: %v", err)
        }
        response := &pb.LogResponse{
               LogFile: buffer[:content], 
        }
        if err := stream.Send(response); err != nil {
            return fmt.Errorf("error sending chunk: %v", err)
        }
    }
	
	return nil 

}

