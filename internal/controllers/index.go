package controllers

import (
	"html/template"
	"net/http"
)

type HomeController struct {
	*BaseController
}

func NewHomeController(base *BaseController) *HomeController {
	return &HomeController{BaseController: base}
}

func (c *HomeController) ServeHandle(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("internal/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
