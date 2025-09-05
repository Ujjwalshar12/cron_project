package main

import (
	"log"
	"net"
	"report/controller"
	"report/cronjob"
	"report/example.com/report/proto"
	"report/model"

	"google.golang.org/grpc"
)

func main() {
	// Create model (state holder)
	reportModel := model.NewReportServiceServer()

	// Create gRPC server and bind
	grpcServer := grpc.NewServer()
	proto.RegisterReportServiceServer(grpcServer, &controller.Server{ReportServiceServer: reportModel})

	// Start cron
	go cronjob.StartCron(reportModel)

	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server started on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
