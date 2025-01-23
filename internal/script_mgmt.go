package internal

import (
	"context"
	"log"
	pb "oob-connector/api/script_mgmt"
	"os/exec"
	"github.com/robfig/cron/v3"
	
)
type ScriptMgmtServer struct {
	pb.UnimplementedScript_MgmtServer
	
}
type GoScriptJob struct {
	FilePath string   
}

func (g GoScriptJob) Run() { 
	command := exec.Command("bash", g.FilePath )
	output, err := command.CombinedOutput()
	if err != nil {
		log.Printf("Error executing Go script: %s, Error: %v\n", g.FilePath, err)
		
	}
	log.Printf("Output from Go script %s:\n%s\n", g.FilePath, string(output))
} // implementation of the run method of cron.job interface
func (s *ScriptMgmtServer) ScheduleScript (context context.Context , request *pb.ScheduleRequest) (*pb.ScheduleResponse , error) {
	job := GoScriptJob{FilePath: request.ScriptPath}
	c := cron.New(cron.WithSeconds())
	id , err := c.AddJob(request.Time ,job)
	if err != nil {
		return &pb.ScheduleResponse{Message: "Error scheduling Job " } , err
	}
	return &pb.ScheduleResponse{Message: "Job scheduled successfully" , Id:int32(id) } ,nil

	 
}