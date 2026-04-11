package domain

import "time"

type URL struct {
	ID         int64
	ShortCode  string
	LongURL    string
	CreatedAt  time.Time
	ExpiresAt  *time.Time
	ClickCount int64
}
