package router

import (
	"encoding/json"
	"goApi/controller"
	"log"
	"net/http"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	//Fazendo a busca pela lista
	employees, err := controller.FindAll()
	if err != nil {
		log.Printf("Erro na requisicao: %v", err)
		http.Error(w, "Erro na requisicao", http.StatusInternalServerError)
		return
	}

	//Retornando o status e a lista
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
