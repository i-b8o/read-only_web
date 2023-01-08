package service

import (
	"context"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/searcher/v1"
)

type SearchStorage interface {
	Search(ctx context.Context, req pb.SearchRequest) ([]*pb.SearchResponse, error)
}

type searchService struct {
	storage SearchStorage
	logger  logging.Logger
}

func NewSearchService(storage SearchStorage, logger logging.Logger) *searchService {
	return &searchService{storage: storage, logger: logger}
}

func (ss searchService) Search(ctx context.Context, searchQuery string, subj pb.SearchRequest_Subject, params ...uint32) ([]*pb.SearchResponse, error) {
	if len(params) != 2 {
		req := pb.SearchRequest{SearchQuery: searchQuery, Subject: subj}
		return ss.storage.Search(ctx, req)
	}
	offset := params[0]
	limit := params[1]
	req := pb.SearchRequest{SearchQuery: searchQuery, Limit: limit, Offset: offset, Subject: subj}
	return ss.storage.Search(ctx, req)
}
