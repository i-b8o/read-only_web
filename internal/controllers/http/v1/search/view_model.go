package search_controller

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/i-b8o/read-only_contracts/pb/searcher/v1"
)

type SearchService interface {
	Search(ctx context.Context, searchQuery string, params ...uint32) ([]*pb.SearchResponse, error)
	DocSearch(ctx context.Context, searchQuery string, params ...uint32) ([]*pb.SearchResponse, error)
	ChSearch(ctx context.Context, searchQuery string, params ...uint32) ([]*pb.SearchResponse, error)
	PSearch(ctx context.Context, searchQuery string, params ...uint32) ([]*pb.SearchResponse, error)
}

type viewModelState struct {
	status int
	resp   []*pb.SearchResponse
}

type viewModel struct {
	searchService SearchService
}

func NewViewModel(searchService SearchService) *viewModel {
	return &viewModel{searchService: searchService}
}

type req struct {
	query  string
	limit  string
	offset string
}

func (r *req) CastParams() (uint32, uint32) {
	// type-conversion then validate id is a positive num
	offset, err := strconv.ParseUint(r.offset, 10, 64)
	if err != nil {
		return 0, 0
	}
	limit, err := strconv.ParseUint(r.limit, 10, 64)
	if err != nil {
		return 0, 0
	}
	return uint32(offset), uint32(limit)
}

func (vm viewModel) GetState(ctx context.Context, r req) *viewModelState {
	offset, limit := r.CastParams()
	if offset == 0 {
		return &viewModelState{status: http.StatusBadRequest}
	}
	resp, err := vm.searchService.Search(ctx, r.query, offset, limit)
	if err != nil {
		return &viewModelState{status: http.StatusBadRequest}
	}
	s := viewModelState{
		status: http.StatusOK,
		resp:   resp,
	}
	return &s
}

func (vm viewModel) GetDState(ctx context.Context, r req) *viewModelState {
	offset, limit := r.CastParams()
	if offset == 0 {
		return &viewModelState{status: http.StatusBadRequest}
	}
	resp, err := vm.searchService.DocSearch(ctx, r.query, offset, limit)
	if err != nil {
		return &viewModelState{status: http.StatusBadRequest}
	}
	s := viewModelState{
		status: http.StatusOK,
		resp:   resp,
	}
	return &s
}

func (vm viewModel) GetCState(ctx context.Context, r req) *viewModelState {
	offset, limit := r.CastParams()
	if offset == 0 {
		return &viewModelState{status: http.StatusBadRequest}
	}
	resp, err := vm.searchService.ChSearch(ctx, r.query, offset, limit)
	if err != nil {
		return &viewModelState{status: http.StatusBadRequest}
	}
	s := viewModelState{
		status: http.StatusOK,
		resp:   resp,
	}
	return &s
}

func (vm viewModel) GetPState(ctx context.Context, r req) *viewModelState {
	offset, limit := r.CastParams()
	if offset == 0 {
		return &viewModelState{status: http.StatusBadRequest}
	}
	resp, err := vm.searchService.PSearch(ctx, r.query, offset, limit)
	if err != nil {
		return &viewModelState{status: http.StatusBadRequest}
	}
	s := viewModelState{
		status: http.StatusOK,
		resp:   resp,
	}
	return &s
}
