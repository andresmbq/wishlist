package models

import "time"

type WishlistItem struct {
	ID        string
	UserID    string
	ProductID string
	AddedAt   time.Time
}
