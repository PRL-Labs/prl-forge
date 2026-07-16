package stratum

import (
	"log"
	"time"

	
)

func (s *Session) jobLoop() {
	var lastJobID string
    job := CurrentJob()
     if job != nil {
    lastJobID = job.ID
}
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		<-ticker.C

		job := CurrentJob()
		if job == nil {
			continue
		}

		if job.ID == lastJobID {
			continue
		}

		lastJobID = job.ID

		log.Printf("📦 Sending new job %s", job.ID)

	
if err := SendCurrentJob(s); err != nil {

			log.Printf("notify error: %v", err)
			return
		}
	}
}