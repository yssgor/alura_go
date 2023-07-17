package routes

import (
	"net/http"

	"github.com/yssgor/alura_go/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
