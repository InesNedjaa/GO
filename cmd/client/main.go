package main

import (
	"context"
	"log"
	"oob-connector-proxy/v2/api/monitoring_mgmt"
	"oob-connector-proxy/v2/api/power_mgmt"
	"oob-connector-proxy/v2/api/script_mgmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


const ( server_addr = "localhost:8080")
func main () {

	conn , err := grpc.NewClient(server_addr,  grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("connection failed")

	}
	defer conn.Close()
	client1 := power_mgmt.NewPower_MgmtClient(conn)
	client2 := monitoring_mgmt.NewMonitoringMgmtClient(conn)
	client3 := script_mgmt.NewScript_MgmtClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_ , err = client1.PowerOn(ctx, &power_mgmt.PowerOnDeviceRequest{MacAddr: "4444444"})
	if err != nil {
		log.Fatalf("Error calling Service1: %v", err)
	}
	_ , err = client2.GetMetrics(ctx, &monitoring_mgmt.MetricsRequest{DeviceId: 1})
	if err != nil {
		log.Fatalf("Error calling Service1: %v", err)
	}
	
	_, err = client3.ScheduleScript(ctx, &script_mgmt.ScheduleRequest{Time: "00:00" , ScriptPath: "/script.sh"})
	if err != nil {
		log.Fatalf("Error calling Service1: %v", err)
	}

}