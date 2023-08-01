package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/models"
	"go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Nome1", Historia: "História 1"},
		{ID: 2, Nome: "Nome2", Historia: "História 2"},
	}

	database.ConnectionDB()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
