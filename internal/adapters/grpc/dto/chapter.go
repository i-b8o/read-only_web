package adapters_grpc_dto

import (
	"read-only_web/internal/domain/entity"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

func ChapterFromGetOneChapterResponse(resp *pb.GetOneChapterResponse) (chapter entity.Chapter) {
	var paragraphs []entity.Paragraph
	for _, p := range resp.Paragraphs {
		paragraph := entity.Paragraph{ID: p.ID, Num: uint64(p.Num),HasLinks: p.HasLinks, IsTable: p.IsTable, IsNFT: p.IsNFT, Class: p.Class, Content: p.Content, ChapterID: p.ChapterID}
		paragraphs := append(paragraphs, paragraph)
	}
	return entity.Chapter{ID: resp.ID, Name: resp.Name, Num: resp.Num, RegulationID: resp.RegulationID, OrderNum: uint64(resp.OrderNum), Paragraphs: paragraphs, UpdatedAt: resp.}
}
