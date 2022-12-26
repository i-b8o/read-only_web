package entity

import "time"

type Doc struct {
	Id          uint64        `json:"id,omitempty"`
	Type        string        `json:"type,omitempty"`
	Subtype     string        `json:"subtype,omitempty"`
	Rev         string        `json:"rev,omitempty"`
	Name        string        `json:"name"`
	Title       string        `json:"title"`
	Keywords    string        `json:"keywords"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time    `json:"updated_at,omitempty"`
	Chapters    []ChapterInfo `json:"chapters"`
}

// TODO Add all docs page
// TODO Add rev section in templates
