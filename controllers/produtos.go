package controllers

import (
	"html/template"
	"net/http"

	"github.com/yssgor/alura_go/models"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}
