package api

import (
	"encoding/json"
	"net/http"
	"wishlist/internal/app"
)

type WishlistHandler struct {
	Service *app.WishlistService
}

func NewWishlistHandler(service *app.WishlistService) *WishlistHandler {
	return &WishlistHandler{
		Service: service,
	}
}

func (h *WishlistHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var req AddItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := h.Service.AddItem(r.Context(), req.UserID, req.ProductID)
	if err != nil {
		http.Error(w, "could not add item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
