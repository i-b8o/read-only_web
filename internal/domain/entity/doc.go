package entity

import "time"

type Doc struct {
	Id          uint64        `json:"id,omitempty"`
	Pseudo      string        `json:"r_pseudo,omitempty"`
	Name        string        `json:"name"`
	Title       string        `json:"title"`
	Keywords    string        `json:"keywords"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time    `json:"updated_at,omitempty"`
	Chapters    []ChapterInfo `json:"chapters"`
}
