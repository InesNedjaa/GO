package internal
import(
	"context"
	"fmt"
	pb "oob-connector/api/power_mgmt"
	"os/exec"
	
)
type PowerMgmtServer struct {
	pb.UnimplementedPower_MgmtServer
	
	
}
func (s *PowerMgmtServer) PowerOn(context context.Context, request *pb.PowerOnDeviceRequest) (*pb.PowerDeviceResponse, error) {

    command := fmt.Sprintf("wakeonlan %s", request.MacAddr)
    cmd := exec.Command("bash", "-c", command)
    _, err := cmd.CombinedOutput()
    if err != nil {
        return nil, err
    }
    return &pb.PowerDeviceResponse{
        Message: "device powered on successfully",}, nil
}

func ( s *PowerMgmtServer) PowerOff (context context.Context , request *pb.PowerDeviceRequest)(*pb.PowerDeviceResponse , error) {
	command := fmt.Sprintf("ssh -i %s %s@%s sudo shutdown -h now", request.PrivateKeyPath, request.Username, request.Host)
    cmd := exec.Command("bash", "-c", command)
	_, err := cmd.CombinedOutput()
	if err != nil {
		 return nil, err
	}
	return &pb.PowerDeviceResponse{Message: "device powered of successfully"} , nil
}

func ( s *PowerMgmtServer) Rebot (context context.Context , request *pb.PowerDeviceRequest)(*pb.PowerDeviceResponse , error) {
	command := fmt.Sprintf("ssh -i %s %s@%s sudo reboot", request.PrivateKeyPath, request.Username, request.Host)
    cmd := exec.Command("bash", "-c", command)
	_, err := cmd.CombinedOutput()
	if err != nil {
		 return nil , err
	}
	return &pb.PowerDeviceResponse{Message: "device rebooted successfully"} , nil
}