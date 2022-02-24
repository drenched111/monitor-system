package main

import (
	"context"
	"flag"
	"google.golang.org/grpc/credentials/insecure"
	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"time"

	"google.golang.org/grpc"
	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "Go_project/report"
)

const (
	defaultName = "world"
)

type report_struct struct {
	Timestamp  int64
	Metric     string
	Dimensions map[string]string
	Value      float64
}

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
	t    report_struct
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewReportServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t.Timestamp = 10000
	t.Metric = "cpu"
	t.Dimensions = make(map[string]string)
	t.Dimensions["ip"] = "1.1.1.1"
	t.Value = 0.9

	r, err := c.Report(ctx, &pb.ReportReq{Timestamp: t.Timestamp, Metric: t.Metric, Dimensions: t.Dimensions, Value: t.Value})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMsg())
	log.Printf("Greeting: %d", r.GetCode())

}
