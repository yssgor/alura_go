package main

import (
	"net/http"

	"github.com/yssgor/alura_go/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":3000", nil)
}
