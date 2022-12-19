package chapter_controller

import "read-only_web/internal/domain/entity"

type ViewModel struct {
	Regulations  []entity.Regulation
	ChapterID    uint64
	Name         string
	Abbreviation string
	Header       *string
	Title        *string
	Meta         *string
	Keywords     *string
	Prev         entity.Chapter
	Next         entity.Chapter
	Num          string
	Paragraphs   []TemplParagraph
	Chapters     []entity.Chapter
	Updated      string
}
