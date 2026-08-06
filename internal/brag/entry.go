package brag

import "time"

type Category string

const (
	Feedback  Category = "feedback"
	Technical Category = "technical"
	Project   Category = "project"
)

type Entry struct {
	Category   Category  `json:"category,omitempty"`
	CreatedAt  time.Time `json:"created,omitempty"`
	ModifiedAt time.Time `json:"modified,omitempty"`
	Value      string    `json:"value,omitempty"`
	Role       string    `json:"role,omitempty"`
	Source     string    `json:"source,omitempty"`
}
