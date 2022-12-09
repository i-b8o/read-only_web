package dto

type CreateParagraphsRequest struct {
	Paragraphs []Paragraph `json:"paragraphs,omitempty"`
}

type Paragraph struct {
	ParagraphID       uint64 `json:"paragraph_id,omitempty"`
	ParagraphOrderNum uint64 `json:"paragraph_order_num"`
	IsTable           bool   `json:"is_table"`
	IsNFT             bool   `json:"is_nft"`
	HasLinks          bool   `json:"has_links"`
	ParagraphClass    string `json:"paragraph_class,omitempty"`
	ParagraphText     string `json:"paragraph_text"`
	ChapterID         uint64 `json:"chapter_id"`
}

func (dto *Paragraph) Validate() {
	if dto.ParagraphClass == "" {
		dto.ParagraphClass = "-"
	}
}
