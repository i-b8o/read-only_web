package templmanager

import (
	"context"
	"errors"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var templates map[string]*template.Template

type TemplateConfig struct {
	TemplatePath string
}

type TemplateManager struct {
	templatePath string
	mainTemplate string
}

func NewTemplateManager(templatePath, mainTemplate string) TemplateManager {
	return TemplateManager{templatePath: templatePath, mainTemplate: mainTemplate}
}

func (tm TemplateManager) LoadTemplates(ctx context.Context) (err error) {
	if tm.templatePath == "nil" {
		return errors.New("TemplateConfig not initialized")
	}
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(tm.mainTemplate)
	if err != nil {
		return err
	}

	curdir, _ := os.Getwd()
	folders, err := OSReadDir(curdir + tm.templatePath)
	if err != nil {
		return err
	}

	layoutFiles, err := filepath.Glob(curdir + tm.templatePath + "/layouts/*.tmpl")
	if err != nil {
		return err
	}

	for _, folder := range folders {
		if folder == "layouts" {
			continue
		}
		includeFiles, err := filepath.Glob(curdir + tm.templatePath + folder + "/*.tmpl")
		if err != nil {
			return err
		}
		folderName := filepath.Base(folder)
		templates[folderName], err = mainTemplate.Clone()
		if err != nil {
			return err
		}
		includeFiles = append(includeFiles, layoutFiles...)
		templates[folderName] = template.Must(templates[folderName].ParseFiles(includeFiles...))
	}
	return nil
}

func (tm TemplateManager) RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template"+name+" does not exist.", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
