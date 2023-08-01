package routes

import (
	"go-api-rest/controllers"
	"go-api-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.OnePersonaidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreateNewPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
