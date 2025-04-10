package internal
import(
	"context"
	"fmt"
	pb "grpc-gateway/api/service2"
	"os/exec"
	
)
type Service2Server struct {
	pb.UnimplementedService2Server
	
	
}
func (s *Service2Server) PowerOn(context context.Context, request *pb.PowerOnDeviceRequest) (*pb.PowerDeviceResponse, error) {

    command := fmt.Sprintf("wakeonlan %s", request.MacAddr)
    cmd := exec.Command("bash", "-c", command) 
    _, err := cmd.CombinedOutput()
    if err != nil {
        return nil, err
    }
    return &pb.PowerDeviceResponse{
        Message: "device powered on successfully",}, nil
}

func ( s *Service2Server) PowerOff (context context.Context , request *pb.PowerDeviceRequest)(*pb.PowerDeviceResponse , error) {
	command := fmt.Sprintf("ssh -i %s %s@%s sudo shutdown -h now", request.PrivateKeyPath, request.Username, request.Host)
    cmd := exec.Command("bash", "-c", command)
	_, err := cmd.CombinedOutput()
	if err != nil {
		 return nil, err
	}
	return &pb.PowerDeviceResponse{Message: "device powered of successfully"} , nil
}

func ( s *Service2Server) Rebot (context context.Context , request *pb.PowerDeviceRequest)(*pb.PowerDeviceResponse , error) {
	command := fmt.Sprintf("ssh -i %s %s@%s sudo reboot", request.PrivateKeyPath, request.Username, request.Host)
    cmd := exec.Command("bash", "-c", command)
	_, err := cmd.CombinedOutput()
	if err != nil {
		 return nil , err
	}
	return &pb.PowerDeviceResponse{Message: "device rebooted successfully"} , nil
}