package app

import (
	"context"
	"wishlist/pkg/models"
)

type WishlistRepository interface {
	Save(ctx context.Context, item *models.WishlistItem) error
	GetByUserID(ctx context.Context, userID string) ([]*models.WishlistItem, error)
	Remove(ctx context.Context, itemID string) error
}
