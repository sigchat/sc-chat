package dto

import "time"

type MessageDTO struct {
	ID            int
	RoomID        int
	MessageStatus string
	UserID        int
	MessageText   *string
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
