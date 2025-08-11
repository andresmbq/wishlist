package app

import (
	"context"
	"errors"
	"testing"
	"time"
	"wishlist/pkg/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// 🧪 Mock Repository
type mockRepo struct {
	SaveFn        func(ctx context.Context, item *models.WishlistItem) error
	GetByUserIDFn func(ctx context.Context, userID string) ([]*models.WishlistItem, error)
	RemoveFn      func(ctx context.Context, itemID string) error
}

func (m *mockRepo) Save(ctx context.Context, item *models.WishlistItem) error {
	return m.SaveFn(ctx, item)
}
func (m *mockRepo) GetByUserID(ctx context.Context, userID string) ([]*models.WishlistItem, error) {
	return m.GetByUserIDFn(ctx, userID)
}
func (m *mockRepo) Remove(ctx context.Context, itemID string) error {
	return m.RemoveFn(ctx, itemID)
}

// 🧪 Mock Publisher
type mockPublisher struct {
	PublishFn func(ctx context.Context, item *models.WishlistItem) error
}

func (m *mockPublisher) PublishWishlistItemAdded(ctx context.Context, item *models.WishlistItem) error {
	return m.PublishFn(ctx, item)
}

func TestAddItem_Success(t *testing.T) {
	repo := &mockRepo{
		SaveFn: func(ctx context.Context, item *models.WishlistItem) error {
			return nil
		},
	}
	publisher := &mockPublisher{
		PublishFn: func(ctx context.Context, item *models.WishlistItem) error {
			return nil
		},
	}

	service := NewWishlistService(repo, publisher)

	item := &models.WishlistItem{
		ID:        uuid.NewString(),
		UserID:    "andres",
		ProductID: "book-123",
		AddedAt:   time.Now(),
	}

	err := service.AddItem(context.Background(), item.UserID, item.ProductID)
	assert.NoError(t, err)
}

func TestAddItem_SaveFails(t *testing.T) {
	repo := &mockRepo{
		SaveFn: func(ctx context.Context, item *models.WishlistItem) error {
			return errors.New("db error")
		},
	}
	publisher := &mockPublisher{
		PublishFn: func(ctx context.Context, item *models.WishlistItem) error {
			return nil
		},
	}

	service := NewWishlistService(repo, publisher)

	item := &models.WishlistItem{
		ID:        uuid.NewString(),
		UserID:    "andres",
		ProductID: "book-123",
		AddedAt:   time.Now(),
	}

	err := service.AddItem(context.Background(), item.UserID, item.ProductID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "db error")
}

func TestGetByUserID(t *testing.T) {
	repo := &mockRepo{
		GetByUserIDFn: func(ctx context.Context, userID string) ([]*models.WishlistItem, error) {
			return []*models.WishlistItem{
				{
					ID:        uuid.NewString(),
					UserID:    userID,
					ProductID: "book-123",
					AddedAt:   time.Now(),
				},
			}, nil
		},
	}

	service := NewWishlistService(repo, nil)

	items, err := service.GetByUserID(context.Background(), "andres")
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, "andres", items[0].UserID)
}

func TestRemoveItem(t *testing.T) {
	repo := &mockRepo{
		RemoveFn: func(ctx context.Context, itemID string) error {
			return nil
		},
	}

	service := NewWishlistService(repo, nil)

	err := service.RemoveItem(context.Background(), "item-123")
	assert.NoError(t, err)
}
