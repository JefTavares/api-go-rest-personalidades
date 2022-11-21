package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeftavares/api-go-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page") //!imprime a msg aonde estiver (no caso localhost:8000)
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
