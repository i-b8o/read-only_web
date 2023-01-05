package search_controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	templateManager "read-only_web/pkg/templmanager"

	"github.com/julienschmidt/httprouter"
)

const (
	search  = "/search"
	rsearch = "/rsearch"
	csearch = "/csearch"
	psearch = "/psearch"
)

type searchHandler struct {
	vm              *viewModel
	templateManager templateManager.TemplateManager
}

func NewDocHandler(vm *viewModel, templateManager templateManager.TemplateManager) *searchHandler {
	return &searchHandler{vm: vm, templateManager: templateManager}
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

	// type-conversion then validate id is a positive num
	uint64Offset, err := strconv.ParseUint(offset, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	uint64Limit, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	searchResults, err := h.searchService.Search(r.Context(), searchQuery, uint32(uint64Offset), uint32(uint64Limit))
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
