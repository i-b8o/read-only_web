package search_controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	templateManager "read-only_web/pkg/templmanager"

	pb_searcher "github.com/i-b8o/read-only_contracts/pb/searcher/v1"
	"github.com/julienschmidt/httprouter"
)

const (
	search = "/search"
)

type searchHandler struct {
	vm              *viewModel
	templateManager templateManager.TemplateManager
}

func NewSearchHandler(vm *viewModel, templateManager templateManager.TemplateManager) *searchHandler {
	return &searchHandler{vm: vm, templateManager: templateManager}
}

func (h *searchHandler) Register(router *httprouter.Router) {
	router.GET(search, h.Search)
}

func parseParams(r *http.Request, params httprouter.Params) *req {
	queryValues := r.URL.Query()
	searchQuery := queryValues.Get("req")
	limit := queryValues.Get("limit")
	offset := queryValues.Get("offset")
	var subj pb_searcher.SearchRequest_Subject
	switch subject := queryValues.Get("subj"); subject {
	case "g":
		subj = pb_searcher.SearchRequest_General
	case "d":
		subj = pb_searcher.SearchRequest_Docs
	case "c":
		subj = pb_searcher.SearchRequest_Chapters
	case "p":
		subj = pb_searcher.SearchRequest_Pargaraphs
	default:
		return nil
	}
	fmt.Printf("query: %s offset: %s, limit: %s, subj: %s", searchQuery, offset, limit, subj)
	return &req{query: searchQuery, offset: offset, limit: limit, subj: subj}
}

func (h *searchHandler) Search(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := parseParams(r, params)
	if req == nil {
		http.Redirect(w, r, "/404", http.StatusNotFound)
		return
	}
	state := h.vm.GetState(r.Context(), req)
	if state == nil {
		http.Redirect(w, r, "/404", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(state.resp)
}
