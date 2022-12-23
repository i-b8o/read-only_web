package entity

import "time"

type Chapter struct {
	ID          uint64      `json:"id,omitempty"`
	Name        string      `json:"name"`
	Num         string      `json:"num,omitempty"`
	DocID       uint64      `json:"doc_id,omitempty"`
	OrderNum    uint32      `json:"order_num"`
	Header      *string     `json:"header"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Keywords    string      `json:"keywords"`
	Paragraphs  []Paragraph `json:"paragraphs,omitempty"`
	UpdatedAt   *time.Time  `json:"updated_at,omitempty"`
}
