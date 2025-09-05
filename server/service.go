package server

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"

	"report/example.com/report/proto"
	"report/model"
)

func GenerateReport(s *model.ReportServiceServer, ctx context.Context, req *proto.GenerateReportRequest) (*proto.GenerateReportResponse, error) {
	reportID := uuid.New().String() //uuid generate

	s.Mu.Lock()
	s.Reports[reportID] = req.UserId //avoid race condition
	s.Mu.Unlock()

	log.Printf("[GenerateReport] UserID=%s ReportID=%s at %s", req.UserId, reportID, time.Now())

	return &proto.GenerateReportResponse{ReportId: reportID}, nil
}

func HealthCheck(ctx context.Context, req *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	log.Printf("[HealthCheck] at %s", time.Now())
	return &proto.HealthCheckResponse{Status: "SERVING"}, nil
}
