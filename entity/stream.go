package entity

import "time"

type StreamSequenceInitTable struct {
	ID            int       `json:"id"`
	Keyword       string    `json:"keyword"`
	Media         string    `json:"media"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	UserAccountID string    `json:"user_account_id"`
}
