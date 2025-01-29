package main

import (
	"fmt"
	"log"
	"net"
	"oob-connector-proxy/v2/api/monitoring_mgmt"
	"oob-connector-proxy/v2/api/power_mgmt"
	"oob-connector-proxy/v2/api/script_mgmt"
	internal "oob-connector-proxy/v2/internal/services"

	"google.golang.org/grpc"
)
const (
	DefaultPort = "8080"
)

 

func StartServer () error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", DefaultPort))
	if err != nil {
		log.Printf("failed to listen on port %s err : %s", DefaultPort, err)
		return err
	}
	grpc := grpc.NewServer()
	power_mgmt.RegisterPower_MgmtServer(grpc, &internal.PowerMgmtServer{})
	script_mgmt.RegisterScript_MgmtServer(grpc, &internal.ScriptMgmtServer{})
	monitoring_mgmt.RegisterMonitoringMgmtServer(grpc, &internal.MonitoringMgmtServer{})
	log.Printf("Server is running on port %s", DefaultPort)
	if err := grpc.Serve(listener); err != nil {
		log.Printf("error while serving : %s", err)
		return err
	}
	return nil

}

func main () {
	
	if err := StartServer(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} 
}