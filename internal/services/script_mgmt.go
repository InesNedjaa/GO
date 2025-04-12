package internal

import (
	"context"
	"log"
	pb "go-proxy/api/script_mgmt"
	"go-proxy/internal/proxy"
)
type ScriptMgmtServer struct {
	pb.UnimplementedScript_MgmtServer
	
}
type GoScriptJob struct {
	FilePath string   
}


func (s *ScriptMgmtServer) ScheduleScript (context context.Context , request *pb.ScheduleRequest) (*pb.ScheduleResponse , error) {
	
	data:= map[string]string{
		"script_path": request.ScriptPath,
		"time": request.Time, 
		}
	oob_module_name:= "module_3"
	err := proxy.ForwardRequest(data,oob_module_name )
	if err != nil {
		log.Fatalf("Error forwarding request: %v", err)
		return nil ,err
	}
	    return nil, nil

	 
}