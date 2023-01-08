package search_provider

import (
	"context"

	pb "github.com/i-b8o/read-only_contracts/pb/searcher/v1"
)

type searchStorage struct {
	client pb.SearcherGRPCClient
}

func NewSearchStorage(client pb.SearcherGRPCClient) *searchStorage {
	return &searchStorage{client: client}
}

func (s *searchStorage) Search(ctx context.Context, req pb.SearchRequest) ([]*pb.SearchResponse, error) {
	resp, err := s.client.Search(ctx, &req)
	if err != nil {
		return nil, err
	}
	return resp.Response, nil
}
