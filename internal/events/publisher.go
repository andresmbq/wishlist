package events

import (
	"context"
	"wishlist/pkg/models"
)

type EventPublisher interface {
	PublishWishlistItemAdded(ctx context.Context, item *models.WishlistItem) error
}
