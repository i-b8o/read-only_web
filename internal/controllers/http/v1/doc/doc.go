package doc_controller

import (
	"net/http"
	templateManager "read-only_web/pkg/templmanager"

	"github.com/julienschmidt/httprouter"
)

const (
	home         = "/"
	reader       = "/reader"
	documentRoot = "/doc/:id"
)

type DocTemplateManager interface {
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
	router.GET(reader, h.Reader)
}

func (h *docHandler) Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	state := h.vm.GetDefaultState()
	h.templateManager.RenderTemplate(w, "home", state)
}

func (h *docHandler) Reader(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	state := h.vm.GetDefaultState()
	h.templateManager.RenderTemplate(w, "reader", state)
}

func (h *docHandler) DocumentRoot(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	state := h.vm.GetState(r.Context(), params.ByName("id"))
	if state == nil {
		http.Redirect(w, r, "/404", http.StatusNotFound)
		return
	}
	h.templateManager.RenderTemplate(w, "doc", state)
}
