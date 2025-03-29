package router

import (
	"encoding/json"
	"goApi/controller"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func FindOne(w http.ResponseWriter, r *http.Request) {
	//Pegando e testando o id passado na requisição
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro na requisicao: %v", err)
		http.Error(w, "Erro na requisicao", http.StatusInternalServerError)
		return
	}

	//Fazendo a pesquisa
	employee, err := controller.FindOne(int64(id))
	if err != nil {
		log.Println("Registro não encontrado")
		http.Error(w, "Registro não encontrado", http.StatusBadRequest)
		return
	}

	//Retornando o sattus e o item
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}
