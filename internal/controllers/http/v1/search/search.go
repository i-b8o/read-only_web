package search_controller

import (
	"context"
	"encoding/json"
	"net/http"
	"read-only_web/internal/domain/entity"

	"github.com/julienschmidt/httprouter"
)

const (
	search  = "/search"
	rsearch = "/rsearch"
	csearch = "/csearch"
	psearch = "/psearch"
)

type SearchService interface {
	Search(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
	RegSearch(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
	ChSearch(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
	PSearch(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
}

type searchHandler struct {
	searchService SearchService
}

func NewSearchHandler(searchService SearchService) *searchHandler {
	return &searchHandler{searchService: searchService}
}

func (h *searchHandler) Register(router *httprouter.Router) {
	router.GET(search, h.Search)
	router.GET(rsearch, h.RSearch)
	router.GET(csearch, h.ChSearch)
	router.GET(psearch, h.PSearch)
}

func (h *searchHandler) Search(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	queryValues := r.URL.Query()
	searchQuery := queryValues.Get("req")

	limit := queryValues.Get("limit")
	offset := queryValues.Get("offset")
	if len(limit) == 0 || len(offset) == 0 {

		searchResults, err := h.searchService.Search(r.Context(), searchQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}
		json.NewEncoder(w).Encode(searchResults)
		return
	}

	searchResults, err := h.searchService.Search(r.Context(), searchQuery, offset, limit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(searchResults)
}

func (h *searchHandler) RSearch(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	queryValues := r.URL.Query()
	searchQuery := queryValues.Get("req")

	limit := queryValues.Get("limit")
	offset := queryValues.Get("offset")
	if len(limit) == 0 || len(offset) == 0 {
		searchResults, err := h.searchService.RegSearch(r.Context(), searchQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}
		json.NewEncoder(w).Encode(searchResults)
		return
	}

	searchResults, err := h.searchService.RegSearch(r.Context(), searchQuery, offset, limit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(searchResults)
}

func (h *searchHandler) ChSearch(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	queryValues := r.URL.Query()
	searchQuery := queryValues.Get("req")

	limit := queryValues.Get("limit")
	offset := queryValues.Get("offset")
	if len(limit) == 0 || len(offset) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("limit and offset must be more than zero")
		searchResults, err := h.searchService.ChSearch(r.Context(), searchQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}
		json.NewEncoder(w).Encode(searchResults)
		return
	}
	searchResults, err := h.searchService.ChSearch(r.Context(), searchQuery, offset, limit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(searchResults)
}

func (h *searchHandler) PSearch(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	queryValues := r.URL.Query()
	searchQuery := queryValues.Get("req")

	limit := queryValues.Get("limit")
	offset := queryValues.Get("offset")
	if len(limit) == 0 || len(offset) == 0 {
		searchResults, err := h.searchService.PSearch(r.Context(), searchQuery, offset, limit)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}
		json.NewEncoder(w).Encode(searchResults)
		return
	}
	searchResults, err := h.searchService.PSearch(r.Context(), searchQuery, offset, limit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(searchResults)
}
