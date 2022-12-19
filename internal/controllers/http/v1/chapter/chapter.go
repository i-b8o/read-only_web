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
	state := h.vm.GetState(r.Context(), params.ByName("id"))
	h.templateManager.RenderTemplate(w, "chapter", state)
}
