package handlers

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"node-week-01-with-chi/store"
	"node-week-01-with-chi/utils"
	"strings"
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
