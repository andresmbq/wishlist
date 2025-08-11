package db

import (
	"context"
	"wishlist/pkg/models"

	"github.com/jackc/pgx/v4"
)

type PostgresWishListRepo struct {
	Conn *pgx.Conn
}

func NewPostgresWishlistRepo(conn *pgx.Conn) *PostgresWishListRepo {
	return &PostgresWishListRepo{Conn: conn}
}

func (r *PostgresWishListRepo) Save(ctx context.Context, item *models.WishlistItem) error {
	_, err := r.Conn.Exec(ctx,
		`INSERT INTO wishlist_items (id, user_id, product_id, created_at)
		 VALUES ($1, $2, $3, $4)`,
		item.ID, item.UserID, item.ProductID, item.AddedAt)
	return err
}

func (r *PostgresWishListRepo) GetByUserID(ctx context.Context, userID string) ([]*models.WishlistItem, error) {
	rows, err := r.Conn.Query(ctx,
		`SELECT id, user_id, product_id, created_at FROM wishlist_items WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.WishlistItem
	for rows.Next() {
		var item models.WishlistItem
		if err := rows.Scan(&item.ID, &item.UserID, &item.ProductID, &item.AddedAt); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func (r *PostgresWishListRepo) Remove(ctx context.Context, itemID string) error {
	_, err := r.Conn.Exec(ctx, `DELETE FROM wishlist_items WHERE id = $1`, itemID)
	return err
}
