package internal

import (
	"context"
	"log"
	pb "grpc-gateway/api/service3"
	"os/exec"
	"github.com/robfig/cron/v3"
	
)
type Service3Server struct {
	pb.UnimplementedService3Server
	
}
type GoScriptJob struct {
	FilePath string   
}

func (g GoScriptJob) Run() { 
	command := exec.Command("bash", g.FilePath )
	output, err := command.CombinedOutput()
	if err != nil {
		log.Printf("Error executing script: %s, Error: %v\n", g.FilePath, err)
		
	}
	log.Printf("Output from script %s:\n%s\n", g.FilePath, string(output))
} // implementation of the run method of cron.job interface
func (s *Service3Server) ScheduleScript (context context.Context , request *pb.ScheduleRequest) (*pb.ScheduleResponse , error) {
	job := GoScriptJob{FilePath: request.ScriptPath}
	c := cron.New(cron.WithSeconds())
	id , err := c.AddJob(request.Time ,job)
	if err != nil {
		return &pb.ScheduleResponse{Message: "Error scheduling Job " } , err
	}
	c.Start()
	return &pb.ScheduleResponse{Message: "Job scheduled successfully" , Id:int32(id) } ,nil

	 
}