package subtypes_controller

import (
	"net/http"
	templateManager "read-only_web/pkg/templmanager"

	"github.com/julienschmidt/httprouter"
)

const (
	subType = "/sub/:type/:subtype"
)

type SubTypesTemplateManager interface {
	RenderTemplate(w http.ResponseWriter, name string, data interface{})
}

type subTypesHandler struct {
	vm              *viewModel
	templateManager templateManager.TemplateManager
}

func NewSubTypesHandler(vm *viewModel, templateManager templateManager.TemplateManager) *subTypesHandler {
	return &subTypesHandler{vm: vm, templateManager: templateManager}
}

func (h *subTypesHandler) Register(router *httprouter.Router) {
	router.GET(subType, h.SubType)
}

func (h *subTypesHandler) SubType(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	state := h.vm.GetState(r.Context(), params.ByName("type"), params.ByName("subtype"))
	if state == nil {
		http.Redirect(w, r, "/404", http.StatusNotFound)
		return
	}
	h.templateManager.RenderTemplate(w, "sub", state)
}
