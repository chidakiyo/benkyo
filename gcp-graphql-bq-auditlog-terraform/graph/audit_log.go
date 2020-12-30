package graph

import (
	"encoding/json"
	"os"
	"time"
)

type AuditLog struct {
	Cookie     string    `json:"cookie"`
	Action     string    `json:"action"`
	ClientIP   string    `json:"client_ip"`
	Timestamp  time.Time `json:"timestamp"`
	Mark       string    `json:"mark"`
	Trace      string    `json:"logging.googleapis.com/trace"`
}

var (
	AuditLogTarget = os.Stdout
)

const Mark = "Target"

func Output(log *AuditLog) {
	log.Mark = Mark // Cloud Logging でフィルタするための条件の type を強制的に入れる
	json.NewEncoder(AuditLogTarget).Encode(log)
}
