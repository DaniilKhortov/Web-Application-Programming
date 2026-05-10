package handler

import (
	"net/http"
	"text/template"
)

// Функція рендеру сторінок-шаблонів  з підставкою данних
func (a *App) render(w http.ResponseWriter, r *http.Request, name string, data any) {

	files := []string{
		"web/templates/" + name + ".html",
		"web/templates/menu.html",
		"web/templates/variants.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		a.serverError(w, r, err)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		a.serverError(w, r, err)
	}
}
