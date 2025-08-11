package events

import "time"

type WishlistItemAddedEvent struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ProductID string    `json:"product_id"`
	AddedAt   time.Time `json:"added_at"`
}
