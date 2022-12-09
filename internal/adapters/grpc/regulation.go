package grpc_adapter

import pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"

type readerStorage struct {
	client pb.ReaderGRPCClient
}

func NewReaderStorage(client pb.ReaderGRPCClient) *readerStorage {
	return &readerStorage{client: client}
}
