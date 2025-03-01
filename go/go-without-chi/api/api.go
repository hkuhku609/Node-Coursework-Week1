package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"node-week-01-without-chi/handlers"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	r := http.NewServeMux()

	quoteHandler := handlers.New()
	r.HandleFunc("GET /", quoteHandler.GetQuotes)
	r.HandleFunc("GET /search", quoteHandler.SearchQuotes)
	r.HandleFunc("GET /random", quoteHandler.RandomQuote)

	quotesRoute := http.NewServeMux()
	quotesRoute.Handle("/quotes/", http.StripPrefix("/quotes", r))

	fmt.Println("Server starting on port 4002")

	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
		RequestJSONMiddleware,
	)

	server := &http.Server{
		Addr: ":4002", Handler: middlewareChain(quotesRoute),
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Error: %v\n", err)
		}
	}()

	fmt.Println("Press Ctrl+C to stop the server")
	<-sigCh

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error:%v/n", err)

	}

	fmt.Println("Server gracefully stopped")

}
func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return handlers.NewLog(next).ServeHTTP
}

func RequestJSONMiddleware(next http.Handler) http.HandlerFunc {
	return handlers.JSONContentTypeMiddleware(next).ServeHTTP
}

func MiddlewareChain(middlewares ...handlers.Middleware) handlers.Middleware {
	return func(next http.Handler) http.HandlerFunc {
		return handlers.MiddlewareHandler(middlewares, next)
	}
}
