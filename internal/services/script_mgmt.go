package internal

import (
	"context"
	pb "oob-connector-proxy/v2/api/script_mgmt"
	"oob-connector-proxy/v2/internal/proxy"
)
type ScriptMgmtServer struct {
	pb.UnimplementedScript_MgmtServer
	
}
type GoScriptJob struct {
	FilePath string   
}


func (s *ScriptMgmtServer) ScheduleScript (context context.Context , request *pb.ScheduleRequest) (*pb.ScheduleResponse , error) {
	
	proxy_request := &proxy.ProxyType{
		Data: map[string]string{
			"script_path": request.ScriptPath,
			"time": request.Time, 
			
		},
		Oob_module: proxy.OobModuleType{
			Name:           "OOB_module_3", 
			Addr:           "127.0.0.1",
			Port:           "8083",
			Connection_type: "IPC",
		},
	}
	proxy.RecieveServiceRequest(proxy_request)
    return nil, nil

	 
}