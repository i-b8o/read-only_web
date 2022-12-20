package doc_controller

import (
	"net/http"
	templateManager "read-only_web/internal/templmanager"

	"github.com/julienschmidt/httprouter"
)

const (
	home         = "/"
	documentRoot = "/doc/:id"
)

type RegTemplateManager interface {
	RenderTemplate(w http.ResponseWriter, name string, data interface{})
}

type docHandler struct {
	vm              *viewModel
	templateManager templateManager.TemplateManager
}

func NewDocHandler(vm *viewModel, templateManager templateManager.TemplateManager) *docHandler {
	return &docHandler{vm: vm, templateManager: templateManager}
}

func (h *docHandler) Register(router *httprouter.Router) {
	router.GET(documentRoot, h.DocumentRoot)
	router.GET(home, h.Home)
}

func (h *docHandler) Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.templateManager.RenderTemplate(w, "home", nil)
}

func (h *docHandler) DocumentRoot(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	state := h.vm.GetState(r.Context(), params.ByName("id"))
	if state == nil {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}
	h.templateManager.RenderTemplate(w, "doc", state)
}
