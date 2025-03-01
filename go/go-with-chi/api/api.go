package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"node-week-01-with-chi/handlers"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	quoteHandler := handlers.New()
	r.Route("/quotes", func(r chi.Router) {
		r.Get("/", quoteHandler.GetQuotes)
		r.Get("/random", quoteHandler.RandomQuote)
		r.Get("/search", quoteHandler.SearchQuotes)
	})

	fmt.Println("Server starting on port 4001")

	srv := http.Server{
		Addr: ":4001", Handler: r,
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Error: %v\n", err)
		}
	}()

	fmt.Println("Press Ctrl+C to stop the server")
	<-sigCh

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Println("Server gracefully stopped")
}
