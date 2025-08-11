package api

import (
	"encoding/json"
	"log"
	"net/http"
	"wishlist/internal/app"

	"github.com/gorilla/mux"
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
		log.Printf("failed to save wishlist item: %v", err)
		http.Error(w, "could not add item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *WishlistHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	if userID == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	items, err := h.Service.Repo.GetByUserID(r.Context(), userID)
	if err != nil {
		log.Printf("error retrieving wishlist for user %s: %v", userID, err)
		http.Error(w, "could not retrieve wishlist", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *WishlistHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["item_id"]
	if itemID == "" {
		http.Error(w, "missing item_id", http.StatusBadRequest)
		return
	}

	err := h.Service.RemoveItem(r.Context(), itemID)
	if err != nil {
		log.Printf("error removing item %s %v", itemID, err)
		http.Error(w, "could not remove item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
