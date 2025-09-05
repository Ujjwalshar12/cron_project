package model

import (
	"report/example.com/report/proto"
	"sync"
)

// ReportServiceServer is the gRPC server implementation
type ReportServiceServer struct {
	proto.UnimplementedReportServiceServer
	Mu      sync.Mutex
	Reports map[string]string // reportID -> userID
}

// Constructor
func NewReportServiceServer() *ReportServiceServer {
	return &ReportServiceServer{
		Reports: make(map[string]string),
	}
}
