package service

import (
	"context"

	pb "github.com/i-b8o/read-only_contracts/pb/searcher/v1"
)

type SearchStorage interface {
	Docs(ctx context.Context, req pb.SearchRequest) ([]*pb.SearchResponse, error)
	Chapters(ctx context.Context, req pb.SearchRequest) ([]*pb.SearchResponse, error)
	Pargaraphs(ctx context.Context, req pb.SearchRequest) ([]*pb.SearchResponse, error)
	General(ctx context.Context, req pb.SearchRequest) ([]*pb.SearchResponse, error)
}

type searchService struct {
	storage SearchStorage
}

func NewSearchService(storage SearchStorage) *searchService {
	return &searchService{storage: storage}
}

func (ss searchService) Search(ctx context.Context, searchQuery string, params ...uint32) ([]*pb.SearchResponse, error) {
	offset := uint32(0)
	limit := uint32(0)
	if len(params) == 2 {
		offset = params[1]
		limit = params[2]
	}
	req := pb.SearchRequest{SearchQuery: searchQuery, Limit: limit, Offset: offset}
	return ss.storage.Docs(ctx, req)
}

func (ss searchService) DocSearch(ctx context.Context, searchQuery string, params ...uint32) ([]*pb.SearchResponse, error) {
	offset := uint32(0)
	limit := uint32(0)
	if len(params) == 2 {
		offset = params[1]
		limit = params[2]
	}
	req := pb.SearchRequest{SearchQuery: searchQuery, Limit: limit, Offset: offset}
	return ss.storage.Chapters(ctx, req)
}

func (ss searchService) ChSearch(ctx context.Context, searchQuery string, params ...uint32) ([]*pb.SearchResponse, error) {
	offset := uint32(0)
	limit := uint32(0)
	if len(params) == 2 {
		offset = params[1]
		limit = params[2]
	}
	req := pb.SearchRequest{SearchQuery: searchQuery, Limit: limit, Offset: offset}
	return ss.storage.Pargaraphs(ctx, req)
}

func (ss searchService) PSearch(ctx context.Context, searchQuery string, params ...uint32) ([]*pb.SearchResponse, error) {
	offset := uint32(0)
	limit := uint32(0)
	if len(params) == 2 {
		offset = params[1]
		limit = params[2]
	}
	req := pb.SearchRequest{SearchQuery: searchQuery, Limit: limit, Offset: offset}
	return ss.storage.General(ctx, req)
}
