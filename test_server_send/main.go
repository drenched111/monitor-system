//send sever

package main

import (
	pb "Go_project/send"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9001, "The server port")
	//manageResp = []*pb.ManageRsp{
	//	&pb.ManageRsp{Ip: "1.1.1.1", Metric: "cpu", Threshold: 0.9, AlertType: 2},
	//	&pb.ManageRsp{Ip: "1.1.1.1", Metric: "cpu", Threshold: 0.8, AlertType: 1},
	//	&pb.ManageRsp{Ip: "1.1.1.1", Metric: "cpu", Threshold: 0.7, AlertType: 0},
	//}
)

type server struct {
	pb.UnimplementedSendServiceServer
}

func (s *server) Send(ctx context.Context, in *pb.SendReq) (*pb.SendRsp, error) {
	//log.Printf("Received: %v", in.GetName())
	resp := &pb.SendRsp{Code: 0, Msg: "send successfully"}
	log.Printf("%d", resp.GetCode())
	log.Printf("%s", resp.GetMsg())
	return resp, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSendServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
