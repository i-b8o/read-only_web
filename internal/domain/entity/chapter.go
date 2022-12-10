package entity

import "time"

type Chapter struct {
	ID           uint64      `json:"id,omitempty"`
	Name         string      `json:"name"`
	Num          string      `json:"num,omitempty"`
	RegulationID uint64      `json:"regulation_id,omitempty"`
	OrderNum     uint32      `json:"order_num"`
	Paragraphs   []Paragraph `json:"paragraphs,omitempty"`
	UpdatedAt    *time.Time  `json:"updated_at,omitempty"`
}

func (c Chapter) IsEmpty() bool {
	if (c.Name == "") && (c.Num == "") && len(c.Paragraphs) == 0 {
		return true
	}
	return false
}
