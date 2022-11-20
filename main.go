package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page") //!imprime a msg aonde estiver (no caso localhost:8000)
}

// lidar com as requisições (tradução livre de handler request ;)
func HandlerRequest() {
	http.HandleFunc("/", Home)                   //cria a rota
	log.Fatal(http.ListenAndServe(":8000", nil)) //log caso de algum erro e sobe o servidor
}

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	HandlerRequest()
}
