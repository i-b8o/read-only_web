package v1

import (
	"context"
	"html/template"
	"net/http"
	"read-only_web/internal/domain/entity"
	templateManager "read-only_web/internal/templmanager"

	"github.com/julienschmidt/httprouter"
)

const (
	chapterGet = "/chapter/:id"
)

type ChapterUsecase interface {
	GetChapter(ctx context.Context, chapterID string) (entity.Regulation, entity.Chapter)
}

type chapterHandler struct {
	chapterUsecase  ChapterUsecase
	templateManager templateManager.TemplateManager
}

func NewChapterHandler(chapterUsecase ChapterUsecase, templateManager templateManager.TemplateManager) *chapterHandler {
	return &chapterHandler{chapterUsecase: chapterUsecase, templateManager: templateManager}
}

func (h *chapterHandler) Register(router *httprouter.Router) {
	router.GET(chapterGet, h.GetChapter)
}

func (h *chapterHandler) GetChapter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	regulation, chapter := h.chapterUsecase.GetChapter(r.Context(), params.ByName("id"))
	// curdir, _ := os.Getwd()

	// templateFiles := []string{curdir + "/internal/templates/root/root.tmpl", curdir + "/internal/templates/root/header.tmpl", curdir + "/internal/templates/chapter/aside.tmpl", curdir + "/internal/templates/chapter/section.tmpl", curdir + "/internal/templates/root/root_css.tmpl", curdir + "/internal/templates/root/mob_root_css.tmpl"}

	// tmp := template.Must(template.ParseFiles(templateFiles...))

	var prevChapter, nextChapter entity.Chapter
	// the chapter order num starts from 1 (not 0)
	if chapter.OrderNum > 1 {
		prevChapter = regulation.Chapters[chapter.OrderNum-2]
	}
	if int(chapter.OrderNum) < len(regulation.Chapters) {
		nextChapter = regulation.Chapters[chapter.OrderNum]
	}

	var templParagraphs []TemplParagraph
	for _, p := range chapter.Paragraphs {
		templParagraph := TemplParagraph{ID: p.ID, Num: p.Num, Class: p.Class, Content: template.HTML(p.Content), ChapterID: p.ChapterID, IsTable: p.IsTable}
		templParagraphs = append(templParagraphs, templParagraph)
	}

	data := Data{
		ChapterID:    chapter.ID,
		Abbreviation: regulation.Abbreviation,
		Name:         chapter.Name,
		Num:          chapter.Num,
		Chapters:     regulation.Chapters,
		Prev:         prevChapter,
		Next:         nextChapter,
		Paragraphs:   templParagraphs,
		Updated:      chapter.UpdatedAt.Format("02.01.2006"),
	}

	h.templateManager.RenderTemplate(w, "chapter", data)
}
