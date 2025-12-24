package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/coelhoedudev/gobit/internal/api"
	"github.com/coelhoedudev/gobit/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatalf("failed to start server: %w", err)
	}

}

func run(ctx context.Context) error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load enviroments variables %s", err.Error())
	}

	pool, err := pgxpool.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	))

	if err != nil {
		return fmt.Errorf("failed to connect to database: %s", err.Error())
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to comunicate with database: %s", err.Error())
	}

	api := api.Api{
		Router:      chi.NewMux(),
		UserService: service.NewUserService(pool),
	}

	api.BindRoutes()
	log.Println(`
		██████╗  ██████╗ 
		██╔════╝ ██╔═══██╗
		██║  ███╗██║   ██║
		██║   ██║██║   ██║
		╚██████╔╝╚██████╔╝
		╚═════╝  ╚═════╝ 
	`)

	fmt.Println("Starting Server on Port :8080")

	server := http.Server{
		Addr:         ":8080",
		Handler:      api.Router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return server.ListenAndServe()

}
