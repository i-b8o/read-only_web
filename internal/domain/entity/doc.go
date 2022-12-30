package entity

type Doc struct {
	ID          uint64        `json:"id,omitempty"`
	Rev         string        `json:"rev,omitempty"`
	Name        string        `json:"name"`
	Title       string        `json:"title"`
	Keywords    string        `json:"keywords"`
	Description string        `json:"description"`
	Chapters    []ChapterInfo `json:"chapters"`
}
