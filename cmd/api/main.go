package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coelhoedudev/gobit/internal/api"
	"github.com/coelhoedudev/gobit/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	ctx := context.Background()
	server, err := createServer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//implementando gracefully shutdown
	//contexto que será cancelado ao receber sigint ou sigterm
	stopCtx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func(ctx context.Context) {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}(ctx)

	<-stopCtx.Done()
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Println("http shutdown error: %w", err)
	}

	log.Println("server stop gracefully")

}

func createServer(ctx context.Context) (*http.Server, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load enviroments variables %s", err.Error())
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
		return nil, fmt.Errorf("failed to connect to database: %s", err.Error())
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to comunicate with database: %s", err.Error())
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

	return &server, nil

}
