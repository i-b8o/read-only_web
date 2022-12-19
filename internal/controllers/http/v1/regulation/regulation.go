package regulation_controller

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

type regulationHandler struct {
	vm              *ViewModel
	templateManager templateManager.TemplateManager
}

func NewRegulationHandler(vm *ViewModel, templateManager templateManager.TemplateManager) *regulationHandler {
	return &regulationHandler{vm: vm, templateManager: templateManager}
}

func (h *regulationHandler) Register(router *httprouter.Router) {
	router.GET(documentRoot, h.DocumentRoot)
	router.GET(home, h.Home)
}

func (h *regulationHandler) Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.templateManager.RenderTemplate(w, "home", nil)
}

func (h *regulationHandler) DocumentRoot(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	state := h.vm.GetState(r.Context(), params.ByName("id"))
	h.templateManager.RenderTemplate(w, "root", state)
}
