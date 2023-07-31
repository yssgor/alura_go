package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/yssgor/alura_go/models"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão", err)
		}

		qtdConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, qtdConvertida)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("nome")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Print("Erro na conversão", err)
		}
		precoconv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Print("Erro na conversão", err)
		}
		quantidadeconv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Print("Erro na conversão", err)
		}

		models.AtualizaProduto(idConv, quantidadeconv, nome, descricao, precoconv)
	}

	http.Redirect(w, r, "/", 301)

}
