package cronjob

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"

	"report/example.com/report/proto"
	"report/server"

	"report/model"
)

// StartCron starts a cron job that generates reports every 10 seconds
func StartCron(srv *model.ReportServiceServer) {
	c := cron.New()
	users := []string{"user1", "user2", "user3"}

	_, err := c.AddFunc("@every 10s", func() {
		for _, u := range users {
			// Call the service layer directly instead of gRPC
			resp, err := server.GenerateReport(srv, context.Background(), &proto.GenerateReportRequest{UserId: u})
			if err != nil {
				log.Printf("[Cron] Error generating report for user=%s: %v", u, err)
				continue
			}
			log.Printf("[Cron] Report generated for user=%s reportID=%s at %s",
				u, resp.ReportId, time.Now().Format(time.RFC3339))
		}
	})
	if err != nil {
		log.Fatalf("Failed to start cron job: %v", err)
	}

	c.Start()
}
