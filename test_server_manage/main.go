//manage sever

package main

import (
	pb "Go_project/manage"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port       = flag.Int("port", 9000, "The server port")
	manageResp = []*pb.ManageRsp{
		&pb.ManageRsp{Ip: "1.1.1.1", Metric: "cpu", Threshold: 0.9, AlertType: 2},
		&pb.ManageRsp{Ip: "1.1.1.1", Metric: "cpu", Threshold: 0.8, AlertType: 1},
		&pb.ManageRsp{Ip: "1.1.1.1", Metric: "cpu", Threshold: 0.7, AlertType: 0},
		//&pb.ManageRsp{},
	}
)

type mserver struct {
	pb.UnimplementedManageServiceServer
}

func (s *mserver) GetThreshold(ctx context.Context, in *pb.ManageReq) (*pb.ManageRspList, error) {
	//log.Printf("Received: %v", in.GetName())
	resp := &pb.ManageRspList{RspList: manageResp}
	//resp := nil
	return resp, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterManageServiceServer(s, &mserver{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
