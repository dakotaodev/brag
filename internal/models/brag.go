package models

import "time"

type Brag struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Category  string    `json:"category"`
	Content   string    `json:"content"`
}
