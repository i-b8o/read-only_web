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

// type TemplParagraph struct {
// 	ID        uint64        `json:"id"`
// 	Num       uint64        `json:"num"`
// 	IsTable   bool          `json:"is_table"`
// 	Class     string        `json:"class,omitempty"`
// 	Content   template.HTML `json:"content,omitempty"`
// 	ChapterID uint64        `json:"chapterid,omitempty"`
// }

// type Data struct {
// 	Regulations  []entity.Regulation
// 	ChapterID    uint64
// 	Name         string
// 	Abbreviation string
// 	Header       *string
// 	Title        *string
// 	Meta         *string
// 	Keywords     *string
// 	Prev         entity.Chapter
// 	Next         entity.Chapter
// 	Num          string
// 	Paragraphs   []TemplParagraph
// 	Chapters     []entity.Chapter
// 	Updated      string
// }

func (h *regulationHandler) Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.templateManager.RenderTemplate(w, "home", nil)
}

func (h *regulationHandler) DocumentRoot(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// TODO get straightaway Data
	// regulation, chapters := h.regulationUsecase.GetDocumentRoot(r.Context(), params.ByName("id"))
	state := h.vm.GetState(r.Context(), params.ByName("id"))
	// if regulation.IsEmpty() || len(chapters) == 0 {
	// w.WriteHeader(404)
	// }
	// data := Data{
	// 	Abbreviation: regulation.Abbreviation,
	// 	Header:       regulation.Header,
	// 	Title:        regulation.Title,
	// 	Meta:         regulation.Meta,
	// 	Keywords:     regulation.Keywords,
	// 	Name:         regulation.Name,
	// 	Chapters:     chapters,
	// }

	h.templateManager.RenderTemplate(w, "root", state)
}