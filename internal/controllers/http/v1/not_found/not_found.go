package not_found_controller

import (
	"net/http"
	templateManager "read-only_web/internal/templmanager"

	"github.com/julienschmidt/httprouter"
)

const (
	notFound = "/404"
)

type notFoundHandler struct {
	templateManager templateManager.TemplateManager
}

func NewNotFoundHandler(templateManager templateManager.TemplateManager) *notFoundHandler {
	return &notFoundHandler{templateManager: templateManager}
}

func (h *notFoundHandler) Register(router *httprouter.Router) {
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
	})
	router.GET(notFound, h.NotFound)

}

func (h *notFoundHandler) NotFound(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	s := getState()
	h.templateManager.RenderTemplate(w, "not_found", s)
}
