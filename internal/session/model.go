package session

import "time"

type session struct{
	UUID string
	UserID int
	CreatedAt time.Time
	ExpiresAt time.Time
}