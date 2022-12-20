package not_found_controller

import (
	"net/http"
	templateManager "read-only_web/internal/templmanager"

	"github.com/julienschmidt/httprouter"
)

type notFoundHandler struct {
	templateManager templateManager.TemplateManager
}

func NewNotFoundHandler(templateManager templateManager.TemplateManager) *notFoundHandler {
	return &notFoundHandler{templateManager: templateManager}
}

func (h *notFoundHandler) Register(router *httprouter.Router) {
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := getState()
		h.templateManager.RenderTemplate(w, "not_found", s)
	})
}
