package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azum, bem bonita", Preco: 39, Quantidade: 4},
		{"Tenis", "Confortavel", 59, 4},
		{"Teste", "Porque não foi", 2, 8},
		{"teste1", "agora vai", 50, 9},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
