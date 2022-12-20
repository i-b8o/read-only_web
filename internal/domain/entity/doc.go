package entity

import "time"

type Doc struct {
	Id           uint64        `json:"id,omitempty"`
	Pseudo       string        `json:"r_pseudo,omitempty"`
	Name         string        `json:"name"`
	Abbreviation string        `json:"abbreviation"`
	Header       *string       `json:"header"`
	Title        *string       `json:"title"`
	Meta         *string       `json:"meta"`
	Keywords     *string       `json:"key_words"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    *time.Time    `json:"updated_at,omitempty"`
	Chapters     []ChapterInfo `json:"chapters"`
}
