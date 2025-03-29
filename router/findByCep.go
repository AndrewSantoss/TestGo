package router

import (
	"encoding/json"
	"goApi/controller"
	"log"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"
)

func FindByCep(w http.ResponseWriter, r *http.Request) {
	//Pegando e testando o CEP passado na requisição
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		log.Print("CEP não informado")
		http.Error(w, "CEP não informado", http.StatusBadRequest)
		return
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(cep) {
		log.Printf("Erro na requisicao, retire o traço do CEP")
		http.Error(w, "Erro na requisicao, retire o traço do CEP", http.StatusBadRequest)
		return
	}

	//Fazendo a pesquisa e testando
	employees, err := controller.FindByCep(cep)
	if err != nil {
		log.Printf("Erro na requisicao: %v", err)
		http.Error(w, "Erro na requisicao", http.StatusInternalServerError)
		return
	}

	//Retornando o status e a lista
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
