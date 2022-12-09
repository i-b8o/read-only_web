package v1

import (
	"context"
	"html/template"
	"net/http"
	"read-only_web/internal/domain/entity"
	templateManager "read-only_web/internal/templmanager"

	"github.com/julienschmidt/httprouter"
)

const (
	home         = "/"
	documentRoot = "/doc/:id"
	documents    = "/docs"
)

type RegulationUsecase interface {
	GetDocumentRoot(ctx context.Context, stringID string) (entity.Regulation, []entity.Chapter)
	GetDocuments(ctx context.Context) []entity.Regulation
}

type RegTemplateManager interface {
	RenderTemplate(w http.ResponseWriter, name string, data interface{})
}

type regulationHandler struct {
	regulationUsecase RegulationUsecase
	templateManager   templateManager.TemplateManager
	useForParsing     bool
}

func NewRegulationHandler(regulationUsecase RegulationUsecase, templateManager templateManager.TemplateManager, useForParsing bool) *regulationHandler {
	return &regulationHandler{regulationUsecase: regulationUsecase, templateManager: templateManager, useForParsing: useForParsing}
}

func (h *regulationHandler) Register(router *httprouter.Router) {
	router.GET(documentRoot, h.DocumentRoot)
	router.GET(documents, h.Documents)
	router.GET(home, h.Home)

}

type TemplParagraph struct {
	ID        uint64        `json:"id"`
	Num       uint64        `json:"num"`
	IsTable   bool          `json:"is_table"`
	Class     string        `json:"class,omitempty"`
	Content   template.HTML `json:"content,omitempty"`
	ChapterID uint64        `json:"chapterid,omitempty"`
}

type Data struct {
	Regulations  []entity.Regulation
	ChapterID    uint64
	Name         string
	Abbreviation string
	Title        string
	Prev         entity.Chapter
	Next         entity.Chapter
	Num          string
	Paragraphs   []TemplParagraph
	Chapters     []entity.Chapter
	Updated      string
}

func (h *regulationHandler) Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.templateManager.RenderTemplate(w, "home", nil)
}

func (h *regulationHandler) DocumentRoot(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	regulation, chapters := h.regulationUsecase.GetDocumentRoot(r.Context(), params.ByName("id"))

	data := Data{
		Abbreviation: regulation.Abbreviation,
		Title:        regulation.Title,
		Name:         regulation.Name,
		Chapters:     chapters,
	}

	h.templateManager.RenderTemplate(w, "root", data)
}

func (h *regulationHandler) Documents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	regulations := h.regulationUsecase.GetDocuments(r.Context())
	data := Data{
		Regulations: regulations,
	}

	h.templateManager.RenderTemplate(w, "docs", data)
}
