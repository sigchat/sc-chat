package model

import "time"

type Message struct {
	ID            int
	RoomID        int
	MessageStatus string
	UserID        int
	MessageText   *string
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
