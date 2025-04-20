package root

import "time"

type Message struct {
	Content    string        `json:"content"`
	RetryCount int           `json:"retryCount"`
	Timestamp  time.Duration `json:"timestamp"`
}
