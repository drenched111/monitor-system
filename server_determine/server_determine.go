package main

import (
	"Go_project/manage"
	pb "Go_project/report"
	"Go_project/send"
	"context"
	"flag"
	"log"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ReportServer struct {
	pb.UnimplementedReportServiceServer
}

var (
	addr_manage = flag.String("addr_manage", "localhost:9000", "the address to connect to")
	addr_send   = flag.String("addr_send", "localhost:9001", "the address to connect to")
	MetricMap   map[int32]string
)

func Manage(ctx context.Context, port *string, req *pb.ReportReq) *manage.ManageRspList {

	// 获取配置
	manageDial, err := grpc.Dial(*port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to connect manage:%v", err)
	}

	client := manage.NewManageServiceClient(manageDial)

	thresholds, err := client.GetThreshold(ctx, &manage.ManageReq{
		Metric: req.GetMetric(),
		Ip:     req.GetDimensions()["ip"],
	})

	if err != nil {
		log.Fatalf("fail to get threshold: %v", err)
	}
	return thresholds
}

func Determine(ctx context.Context, thresholds *manage.ManageRspList, req *pb.ReportReq) {

	MetricMap = make(map[int32]string)
	MetricMap[0] = "WARN"
	MetricMap[1] = "SEVER"
	MetricMap[2] = "FATAL"

	for _, rsp := range thresholds.GetRspList() { // 因为得到的配置信息是按告警类型严重程度排序的，所以可遍历逐一判断

		if req.GetValue() >= rsp.Threshold {
			// 调用发送服务
			Send(ctx, rsp, req, addr_send)
			break

		}
	}

	// err = influx.WriteReport(req) // 写入influxDB
	// if err != nil {
	// 	return nil, err
	// }
}
func Send(ctx context.Context, rsp *manage.ManageRsp, req *pb.ReportReq, port *string) {

	sendDial, err := grpc.Dial(*port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to connect send: %v", err)
	}

	sendServiceClient := send.NewSendServiceClient(sendDial)
	// 调用发送服务
	_, err = sendServiceClient.Send(ctx, &send.SendReq{
		Timestamp:  req.GetTimestamp(),
		Metric:     req.GetMetric(),
		Dimensions: req.GetDimensions(),
		Value:      req.GetValue(),
		AlertType:  MetricMap[rsp.GetAlertType()],
	})

	if err != nil {
		log.Fatalf("fail to send an e-mail:%v", err)
	}
}

func (s *ReportServer) Report(ctx context.Context, req *pb.ReportReq) (*pb.ReportRsp, error) {

	thresholds := Manage(ctx, addr_manage, req) //获取管理模块的配置信息

	Determine(ctx, thresholds, req) //调用判定模块

	return &pb.ReportRsp{Code: 0, Msg: "success"}, nil

}

// ReportServe 启动判定模块的服务端，main方法可直接调用
func ReportServe() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReportServiceServer(s, &ReportServer{})
	log.Printf("rpc listening: %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("fail to listening %v, error: %v", listen.Addr(), err)
	}
}

func main() {
	ReportServe()

}
