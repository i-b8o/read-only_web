package templmanager

import (
	"context"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/i-b8o/logging"
)

var templates map[string]*template.Template

type TemplateConfig struct {
	TemplatePath string
}

type TemplateManager struct {
	templatePath string
}

func NewTemplateManager(templatePath string) TemplateManager {
	return TemplateManager{templatePath: templatePath}
}

func (tm TemplateManager) LoadTemplates(ctx context.Context, logger logging.Logger) (err error) {
	if tm.templatePath == "nil" {
		logger.Errorf("TemplateConfig not initialized")
		return err
	}
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		logger.Errorf(err.Error())
	}

	curdir, _ := os.Getwd()
	logger.Infof("current directory: %s\n", curdir)
	logger.Infof("template path: %s\n", tm.templatePath)
	folders, err := OSReadDir(curdir + tm.templatePath)
	if err != nil {
		logger.Fatal(err)
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
			logger.Fatal(err)
		}
		folderName := filepath.Base(folder)
		templates[folderName], err = mainTemplate.Clone()
		if err != nil {
			return err
		}
		includeFiles = append(includeFiles, layoutFiles...)
		logger.Infof("path: %s\n", folderName)
		templates[folderName] = template.Must(templates[folderName].ParseFiles(includeFiles...))
	}
	logger.Info("templates loaded successfully")
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
