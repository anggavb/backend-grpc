package main

import (
	"context"
	"log"
	"os"

	"github.com/backend-grpc/api"
	db "github.com/backend-grpc/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading env. \ncause: %s", err.Error())
	}

	conn, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("cannot connect to db:", err.Error())
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(os.Getenv("SERVER_ADDRESS")); err != nil {
		log.Fatal("cannot start server:", err.Error())
	}
}
