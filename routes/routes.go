package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeftavares/api-go-rest/controllers"
)

// lidar com as requisições (tradução livre de handler request ;)
func HandlerRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home) //cria a rota
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", r)) //log caso de algum erro e sobe o servidor
}
