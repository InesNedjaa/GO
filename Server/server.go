package main

import (
	context "context"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"

	pb "metrics-ms/proto/metrics"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)



type MetricType struct {
	CPU      float32 `json:"cpu"`
	Memory   float32 `json:"memory"`
	Bandwidth float32 `json:"bandwidth"`
}

type server struct {
	
	pb.UnimplementedMetricsServer
}
var metricList = make(map[int32][]MetricType)
var uid int32 

func (s *server) RegisterDevice (ctx context.Context, device *pb.RegisterDeviceRequest )  (*pb.RegisterDeviceResponse,  error) { 
	log.Printf("device id:%v" , device.GetUid());
	uid = device.GetUid()
	return &pb.RegisterDeviceResponse{Message: "device registered"} , nil
 }  
func (s *server)StreamMetrics (stream grpc.ClientStreamingServer[pb.SendMetricsData, pb.Metricsack]) error { 
	var metrics []MetricType
	for {
		metric , err := stream.Recv()
		log.Printf("cpu usage:%v" , metric.GetCpu());
		if err==io.EOF { 
			
			log.Printf("stream recieved");
			
			metricList[uid] = metrics 
		    log.Print(metricList)
			return stream.SendAndClose(
				&pb.Metricsack{Message: "stream recieved"} ,
			)
		}
		if err != nil {return err} 
        metrics = append(metrics , MetricType{
			CPU: metric.Cpu,
			Memory : metric.Memory,
			Bandwidth:metric.Bandwith ,
		})
		
	}
	
} 
func GetAllMetrics (context *gin.Context) {
	var list []MetricType  
    for _, lst := range metricList{
        list = append(list, lst...)
    }
    context.JSON(http.StatusOK, list)
}
func GetMetricByID(context *gin.Context) {

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	if list, exists := metricList[int32(id)]; exists {
		context.JSON(http.StatusOK, list)
	} 
}
func main() {
	go func () {
		lis , err := net.Listen("tcp" , ":8888") 
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterMetricsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	}()
	
	router := gin.Default()
	router.GET("/allmetrics" , GetAllMetrics)
	router.GET("/metric/:id" , GetMetricByID)
	router.Run("localhost:8080")
  
}
