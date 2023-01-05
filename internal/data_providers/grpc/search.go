package search_provider

import (
	"context"

	pb "github.com/i-b8o/read-only_contracts/pb/searcher/v1"
)

type searchStorage struct {
	client pb.SearcherGRPCClient
}

func NewChapterStorage(client pb.SearcherGRPCClient) *searchStorage {
	return &searchStorage{client: client}
}

func (s *searchStorage) Docs(ctx context.Context, req pb.SearchRequest) ([]*pb.SearchResponse, error) {
	resp, err := s.client.Docs(ctx, &req)
	if err != nil {
		return nil, err
	}
	return resp.Response, nil
}

func (s *searchStorage) Chapters(ctx context.Context, req pb.SearchRequest) ([]*pb.SearchResponse, error) {
	resp, err := s.client.Chapters(ctx, &req)
	if err != nil {
		return nil, err
	}
	return resp.Response, nil
}

func (s *searchStorage) Pargaraphs(ctx context.Context, req pb.SearchRequest) ([]*pb.SearchResponse, error) {
	resp, err := s.client.Pargaraphs(ctx, &req)
	if err != nil {
		return nil, err
	}
	return resp.Response, nil
}

func (s *searchStorage) General(ctx context.Context, req pb.SearchRequest) ([]*pb.SearchResponse, error) {
	resp, err := s.client.General(ctx, &req)
	if err != nil {
		return nil, err
	}
	return resp.Response, nil
}
