package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	rand "math/rand/v2"
	pb "metrics-ms/proto/metrics"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)   
type Config struct {
	ServerAddr string `json:"server_addr"`
	Interval   int    `json:"interval"`
}
func LoadConfig(file string) (*Config) { 

	configFile, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
		return nil
	}
	
	defer configFile.Close()
	var config Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil
	}

	if config.Interval < 10 || config.Interval > 120 {
		return nil
	}
	return &config
}

func randomUID() int32 {
	return rand.Int32()
}
func CPuUsage () float64{
	cpu_usage , err := cpu.Percent(10*time.Second , false)
	if err != nil {
		fmt.Println("Error in getting CPU usage:", err)
	}
    return cpu_usage[0]  
}
func MemoryUsage () int64{
	memory_usage , err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error in getting memory usage:", err)
	}
	return int64 (memory_usage.Used)
}

func Bandwidth () float32{

	initialStats, err := net.IOCounters(false)
	if err != nil {
		fmt.Println("Error fetching initial network stats:", err)
		return 0
	}
	time.Sleep(1 * time.Second)
	finalStats, err := net.IOCounters(false) 
	if err != nil {
		fmt.Println("Error fetching final network stats:", err)
		return 0
	}
	if len(initialStats) == 0 || len(finalStats) == 0 {
		fmt.Println("No network interfaces found.")
		return 0
	}
	bytesSent := finalStats[0].BytesSent - initialStats[0].BytesSent
	bytesRecv := finalStats[0].BytesRecv - initialStats[0].BytesRecv
	totalBandwidth := bytesSent + bytesRecv 
	return float32(totalBandwidth) 


}
func main () {
	config := LoadConfig("Agent/config.json")
	serverAddr := config.ServerAddr 
	
	// connection
	conn , err := grpc.NewClient(serverAddr,  grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("connection failed")

	}
	defer conn.Close()

	client := pb.NewMetricsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()

    // Send user id
	uid := randomUID() ;
	request := &pb.RegisterDeviceRequest{Uid: uid} 
	response , err := client.RegisterDevice(ctx, request)

	if err != nil {
		fmt.Printf("connection failed")

	}
	fmt.Printf("%s", response.Message)


	// Send Metrics
     
	
	stream , err := client.StreamMetrics(context.Background()) 
	

	for i:=1 ; i<4 ; i++ {
		
		 requestt := &pb.SendMetricsData{
			Cpu:float32(CPuUsage()),
            Memory: float32(MemoryUsage()),
			Bandwith: Bandwidth(),

		 }
		 
		 err := stream.Send(requestt) 
		 if err != nil {
			log.Printf("error in sending metrics")
		 }
		 

	}
	stream.CloseAndRecv()
	 

}
