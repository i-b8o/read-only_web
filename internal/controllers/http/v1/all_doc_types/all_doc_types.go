package alldoctypes_controller

import (
	"net/http"
	templateManager "read-only_web/pkg/templmanager"

	"github.com/julienschmidt/httprouter"
)

const (
	allDocTypes = "/all"
)

type AllDocTypesTemplateManager interface {
	RenderTemplate(w http.ResponseWriter, name string, data interface{})
}

type allDocTypesHandler struct {
	vm              *viewModel
	templateManager templateManager.TemplateManager
}

func NewAllDocTypesHandler(vm *viewModel, templateManager templateManager.TemplateManager) *allDocTypesHandler {
	return &allDocTypesHandler{vm: vm, templateManager: templateManager}
}

func (h *allDocTypesHandler) Register(router *httprouter.Router) {
	router.GET(allDocTypes, h.AllDocTypes)
}

func (h *allDocTypesHandler) AllDocTypes(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	state := h.vm.GetState(r.Context())
	if state == nil {
		http.Redirect(w, r, "/404", http.StatusNotFound)
		return
	}
	h.templateManager.RenderTemplate(w, "all_doc_types", state)
}
