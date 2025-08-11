package api

type AddItemRequest struct {
	UserID    string `json:"user_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
}
