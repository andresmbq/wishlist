package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"wishlist/internal/api"
	"wishlist/internal/app"
	"wishlist/internal/db"
	"wishlist/internal/events"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

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
	r.HandleFunc("/wishlist/{user_id}", handler.GetItems).Methods("GET")
	r.HandleFunc("/wishlist/items/{item_id}", handler.RemoveItem).Methods("DELETE")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("wishlist service running on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
