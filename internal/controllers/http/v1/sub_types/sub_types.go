package subtypes_controller

import (
	"net/http"
	templateManager "read-only_web/pkg/templmanager"

	"github.com/julienschmidt/httprouter"
)

const (
	subTypes = "/subs/:id"
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
	router.GET(subTypes, h.SubTypes)
}

func (h *subTypesHandler) SubTypes(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	state := h.vm.GetState(r.Context(), params.ByName("id"))
	if state == nil {
		http.Redirect(w, r, "/404", http.StatusNotFound)
		return
	}
	h.templateManager.RenderTemplate(w, "sub_types", state)
}
