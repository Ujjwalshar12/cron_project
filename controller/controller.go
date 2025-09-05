package controller

import (
	"context"

	"report/example.com/report/proto"
	"report/model"
	"report/server"
)

// Server binds gRPC -> service layer
type Server struct {
	*model.ReportServiceServer
}

// GenerateReport
func (s *Server) GenerateReport(ctx context.Context, req *proto.GenerateReportRequest) (*proto.GenerateReportResponse, error) {
	return server.GenerateReport(s.ReportServiceServer, ctx, req)
}

// HealthCheck
func (s *Server) HealthCheck(ctx context.Context, req *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	return server.HealthCheck(ctx, req)
}
