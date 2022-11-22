package main

import (
	"fmt"

	"github.com/jeftavares/api-go-rest/models"
	"github.com/jeftavares/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome 1", Historia: "historia 1"},
		{Id: 2, Nome: "nome 2", Historia: "historia 2"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandlerRequest()
}
