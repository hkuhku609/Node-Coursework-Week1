package handlers

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"node-week-01-without-chi/store"
	"node-week-01-without-chi/utils"
	"strings"
	"time"
)

type Quote struct {
	Quotes []store.Quote
}

func New() *Quote {
	return &Quote{
		Quotes: []store.Quote{},
	}
}

func (h *Quote) GetQuotes(w http.ResponseWriter, r *http.Request) {
	quotes := loadJSONfileAndSaveOnQuotes(w)

	if writeErr := utils.WriteJSON(w, http.StatusOK, quotes); writeErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, writeErr.Error())
	}
}

func (h *Quote) RandomQuote(w http.ResponseWriter, r *http.Request) {
	quotes := loadJSONfileAndSaveOnQuotes(w)

	if len(quotes) == 0 {
		utils.WriteError(w, http.StatusNotFound, "Not found the quote")
		return
	}

	randomIndex := rand.IntN(len(quotes))
	randomQuote := quotes[randomIndex]

	if writeErr := utils.WriteJSON(w, http.StatusOK, randomQuote); writeErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, writeErr.Error())
	}
}
func (h *Quote) SearchQuotes(w http.ResponseWriter, r *http.Request) {
	quotes := loadJSONfileAndSaveOnQuotes(w)

	term := r.URL.Query().Get("term")
	if term == "" {
		utils.WriteError(w, http.StatusBadRequest, "keyword can not be null")
		return
	}

	var matchedQuotes []store.Quote
	for _, quote := range quotes {
		if strings.Contains(strings.ToLower(quote.Quote), strings.ToLower(term)) ||
			strings.Contains(strings.ToLower(quote.Author), strings.ToLower(term)) {
			matchedQuotes = append(matchedQuotes, quote)
		}
	}

	if len(matchedQuotes) == 0 {
		utils.WriteError(w, http.StatusNotFound, "Not found the quotes that has matched")
		return
	}

	if writeErr := utils.WriteJSON(w, http.StatusOK, matchedQuotes); writeErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, writeErr.Error())
	}
}

func loadJSONfileAndSaveOnQuotes(w http.ResponseWriter) []store.Quote {
	content, readErr := store.ReadJSONFile("./quotes.json")
	if readErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, readErr.Error())
		return []store.Quote{}
	}

	var quotes []store.Quote

	if err :=
		json.Unmarshal(content, &quotes); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return []store.Quote{}
	}

	return quotes
}
func NewLog(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		elapsedTime := time.Since(startTime)
		log.Printf("[%s] [%s] [%s]\n", r.Method, r.URL.Path, elapsedTime)
	}
}

func JSONContentTypeMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}
func SeeLogger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method: %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareHandler(middlewares []Middleware, next http.Handler) http.HandlerFunc { //<------add

	for i := len(middlewares) - 1; i >= 0; i-- {
		next = middlewares[i](next)
	}
	return next.ServeHTTP
}
