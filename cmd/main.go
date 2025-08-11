package main

import (
	"context"
	"log"
	"os"
	"wishlist/internal/api"
	"wishlist/internal/app"
	"wishlist/internal/db"
	"wishlist/internal/events"

	"github.com/gorilla/mux"

	"github.com/jackc/pgx/v4"
)

func main() {
	dsn := os.Getenv("DB_URL")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("could not connect to the db: %v", err)
	}

	repo := db.NewPostgresWishlistRepo(conn)

	brokers := []string{"localhost:9092"}
	topic := "wishlist.item.added"
	publisher := events.NewKafkaPublisher(brokers, topic)
	service := app.NewWishlistService(repo, publisher)

	handler := api.NewWishlistHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/wishlist/items", handler.AddItem).Methods("POST")
	r.HandleFunc("/wishlist/{user_id}", handler.AddItem).Methods("GET")
	r.HandleFunc("/wishlist/items/{item_id}", handler.AddItem).Methods("DELETE")
}
