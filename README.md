#  Go gRPC Report Service with Cron Jobs

This project demonstrates a **Go microservice** that exposes a **gRPC API** for generating reports, uses a **cron job** to periodically generate reports for predefined users, and stores reports in memory.  
The project is organized with a **3-layer architecture**: `model`, `cronjob`, and `server`.

---

## Features

1. gRPC API
   - `GenerateReport(UserID)` → Generates a report ID for a given user and stores it in memory.
   - `HealthCheck()` → Returns service status (`SERVING`).

2. Cron Job
   - Runs every 10 seconds (configurable).
   - Automatically generates reports for predefined users (`user1`, `user2`, `user3`).

3. In-Memory Storage
   - Reports stored as `reportID -> userID` using a mutex for thread safety.

4. Logging
   - All operations (`GenerateReport`, `HealthCheck`, cron jobs) are logged with timestamps.

5. Clean Architecture
   - model → Data structures and state.
   - server → Business logic (report generation, health check).
   - controller → gRPC bindings (delegates calls to service layer).
   - cronjob → Periodic job scheduler.


## Setup & Run

### 1 Install Dependencies
go mod tidy

### 2 Generate gRPC Code 
protoc --go_out=. --go-grpc_out=. proto/report.proto 

### 3 Run Service 
go run main.go
### 4 gRPC calls
1 Health check
  grpcurl -plaintext localhost:50051 report.ReportService/HealthCheck

2 Generate a report for a user
  grpcurl -plaintext -d '{"user_id": "userX"}' localhost:50051 report.ReportService/GenerateReport


