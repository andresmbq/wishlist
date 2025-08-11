package app

import (
	"context"
	"time"
	"wishlist/internal/events"
	"wishlist/pkg/models"

	"github.com/google/uuid"
)

type WishlistService struct {
	Repo      WishlistRepository
	Publisher events.EventPublisher
}

func NewWishlistService(repo WishlistRepository, publisher events.EventPublisher) *WishlistService {
	return &WishlistService{
		Repo:      repo,
		Publisher: publisher,
	}
}

func (s *WishlistService) AddItem(ctx context.Context, userID, productID string) error {
	item := &models.WishlistItem{
		ID:        uuid.NewString(),
		UserID:    userID,
		ProductID: productID,
		AddedAt:   time.Now(),
	}

	if err := s.Repo.Save(ctx, item); err != nil {
		return err
	}

	return s.Publisher.PublishWishlistItemAdded(ctx, item)
}

func (s *WishlistService) GetByUserID(ctx context.Context, userID string) ([]*models.WishlistItem, error) {
	return s.Repo.GetByUserID(ctx, userID)
}

func (s *WishlistService) RemoveItem(ctx context.Context, itemID string) error {
	return s.Repo.Remove(ctx, itemID)
}
