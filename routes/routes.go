package routes

import (
	"log"
	"net/http"

	"github.com/jeftavares/api-go-rest/controllers"
)

// lidar com as requisições (tradução livre de handler request ;)
func HandlerRequest() {
	http.HandleFunc("/", controllers.Home) //cria a rota
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil)) //log caso de algum erro e sobe o servidor
}
