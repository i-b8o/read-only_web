package chapter_controller

import (
	"net/http"
	templateManager "read-only_web/internal/templmanager"

	"github.com/julienschmidt/httprouter"
)

const (
	chapterGet = "/chapter/:id"
)

type chapterHandler struct {
	vm              *viewModel
	templateManager templateManager.TemplateManager
}

func NewChapterHandler(vm *viewModel, templateManager templateManager.TemplateManager) *chapterHandler {
	return &chapterHandler{vm: vm, templateManager: templateManager}
}

func (h *chapterHandler) Register(router *httprouter.Router) {
	router.GET(chapterGet, h.GetChapter)
}

func (h *chapterHandler) GetChapter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// regulation, chapter := h.chapterUsecase.GetChapter(r.Context(), params.ByName("id"))
	// if regulation.IsEmpty() || chapter.IsEmpty() {
	// 	w.WriteHeader(404)
	// 	h.templateManager.RenderTemplate(w, "404", "asdasdasdasda")
	// 	return
	// }
	// var prevChapter, nextChapter entity.Chapter

	// // the chapter order num starts from 1 (not 0)
	// if chapter.OrderNum > 1 {
	// 	prevChapter = regulation.Chapters[chapter.OrderNum-2]
	// }
	// if int(chapter.OrderNum) < len(regulation.Chapters) {
	// 	nextChapter = regulation.Chapters[chapter.OrderNum]
	// }

	// var templParagraphs []TemplParagraph
	// for _, p := range chapter.Paragraphs {
	// 	templParagraph := TemplParagraph{ID: p.ID, Num: p.Num, Class: p.Class, Content: template.HTML(p.Content), ChapterID: p.ChapterID, IsTable: p.IsTable}
	// 	templParagraphs = append(templParagraphs, templParagraph)
	// }

	// data := Data{
	// 	ChapterID:    chapter.ID,
	// 	Abbreviation: regulation.Abbreviation,
	// 	Name:         chapter.Name,
	// 	Num:          chapter.Num,
	// 	Chapters:     regulation.Chapters,
	// 	Prev:         prevChapter,
	// 	Next:         nextChapter,
	// 	Paragraphs:   templParagraphs,
	// 	Updated:      chapter.UpdatedAt.Format("02.01.2006"),
	// }
	state := h.vm.GetState(r.Context(), params.ByName("id"))
	h.templateManager.RenderTemplate(w, "chapter", state)
}
