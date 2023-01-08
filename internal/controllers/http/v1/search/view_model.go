package search_controller

import (
	"context"
	"strconv"

	pb "github.com/i-b8o/read-only_contracts/pb/searcher/v1"
)

type SearchService interface {
	Search(ctx context.Context, searchQuery string, subj pb.SearchRequest_Subject, params ...uint32) ([]*pb.SearchResponse, error)
}

type viewModelState struct {
	resp []*pb.SearchResponse
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
	subj   pb.SearchRequest_Subject
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

func (vm viewModel) GetState(ctx context.Context, r *req) *viewModelState {
	offset, limit := r.CastParams()
	if limit != 0 {
		resp, err := vm.searchService.Search(ctx, r.query, r.subj, offset, limit)
		if err != nil {
			return nil
		}
		s := viewModelState{
			resp: resp,
		}
		return &s

	}
	resp, err := vm.searchService.Search(ctx, r.query, r.subj)
	if err != nil {
		return nil
	}
	s := viewModelState{
		resp: resp,
	}
	return &s

}
