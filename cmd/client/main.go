package main

import (
	"context"
	"log"
	"go-proxy/api/monitoring_mgmt"
	"go-proxy/api/power_mgmt"
	"go-proxy/api/script_mgmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


const ( server_addr = "localhost:8080")
func main () {

	conn , err := grpc.NewClient(server_addr,  grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Connection to grpc server failed")

	}
	defer conn.Close()
	log.Println("Connected to grpc server")
	client1 := power_mgmt.NewPower_MgmtClient(conn)
	client2 := monitoring_mgmt.NewMonitoringMgmtClient(conn)
	client3 := script_mgmt.NewScript_MgmtClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_ , err = client1.PowerOn(ctx, &power_mgmt.PowerOnDeviceRequest{MacAddr: "4444444"})
	if err != nil {
		log.Fatalf("Error calling Power management service: %v", err)
	}
	_ , err = client2.GetMetrics(ctx, &monitoring_mgmt.MetricsRequest{DeviceId: 1})
	if err != nil {
		log.Fatalf("Error calling Monitoring management service: %v", err)
	}
	
	_, err = client3.ScheduleScript(ctx, &script_mgmt.ScheduleRequest{Time: "00:00" , ScriptPath: "/script.sh"})
	if err != nil {
		log.Fatalf("Error calling Script management service: %v", err)
	}

}