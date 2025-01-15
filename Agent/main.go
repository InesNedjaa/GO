package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	pb "oob-connector/proto/sysupdates"
	"os"
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)
type server struct {
	pb.UnimplementedPower_MgmtServer
	pb.UnimplementedScript_MgmtServer
	pb.UnimplementedMonitoringMgmtServer
	
}
type GoScriptJob struct {
	FilePath string   
}

func (s *server) PowerOn(context context.Context, request *pb.PowerOnDeviceRequest) (*pb.PowerDeviceResponse, error) {

    command := fmt.Sprintf("wakeonlan %s", request.MacAddr)
    cmd := exec.Command("bash", "-c", command)
    _, err := cmd.CombinedOutput()
    if err != nil {
        return nil, err
    }
    return &pb.PowerDeviceResponse{
        Message: "device powered on successfully",}, nil
}

func ( s *server) PowerOff (context context.Context , request *pb.PowerDeviceRequest)(*pb.PowerDeviceResponse , error) {
	command := fmt.Sprintf("ssh -i %s %s@%s sudo shutdown -h now", request.PrivateKeyPath, request.Username, request.Host)
    cmd := exec.Command("bash", "-c", command)
	_, err := cmd.CombinedOutput()
	if err != nil {
		 return nil, err
	}
	return &pb.PowerDeviceResponse{Message: "device powered of successfully"} , nil
}

func ( s *server) Rebot (context context.Context , request *pb.PowerDeviceRequest)(*pb.PowerDeviceResponse , error) {
	command := fmt.Sprintf("ssh -i %s %s@%s sudo reboot", request.PrivateKeyPath, request.Username, request.Host)
    cmd := exec.Command("bash", "-c", command)
	_, err := cmd.CombinedOutput()
	if err != nil {
		 return nil , err
	}
	return &pb.PowerDeviceResponse{Message: "device rebooted successfully"} , nil
}
func (g GoScriptJob) Run() { 
	command := exec.Command("go", "run" , g.FilePath )
	output, err := command.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Go script: %s, Error: %v\n", g.FilePath, err)
		
	}
	log.Printf("Output from Go script %s:\n%s\n", g.FilePath, string(output))
} // implementation of the run method of cron.job interface
func (s *server) ScheduleScript (context context.Context , request *pb.ScheduleRequest) (*pb.ScheduleResponse , error) {
	job := GoScriptJob{FilePath: request.ScriptPath}
	c := cron.New(cron.WithSeconds())
	id , err := c.AddJob(request.Time ,job)
	if err != nil {
		return &pb.ScheduleResponse{Message: "Error scheduling Job " } , err
	}
	return &pb.ScheduleResponse{Message: "Job scheduled successfully" , Id:int32(id) } ,nil

	 
}

func (s *server) GetMetrics (context context.Context , request *pb.MetricsRequest) (*pb.MetricsResponse , error) {
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
	// remember to change network bandwidth to uint64
	return &pb.MetricsResponse{CpuUsage: float32(cpu_usage_percentage) , MemoryUsage: float32(memory_usage_percentage), NetworkBandwidth: network_bandwidth} , finalError
	
}

func (s *server) GetLogs (context context.Context , request *pb.LogsRequest) (*pb.LogResponse , error) {
	var file_path string 
    switch request.GetLogType(){
	case pb.LogType_System_logs : file_path="/var/log/syslog"
	case pb.LogType_Kernel_logs : file_path ="/var/log/kern.log"
	case pb.LogType_Auth_logs : file_path ="/var/log/auth.log"
	}
	file, err := os.OpenFile(file_path, os.O_RDONLY , 0)
	defer file.Close()
	if err != nil {
		return nil , err
	}
	file_content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return &pb.LogResponse{LogFile: file_content } ,nil

}

func main () {

	
}