package router

import (
	"encoding/json"
	"goApi/controller"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	//Pegando e testando o id passado na requisição
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro na requisicao: %v", err)
		http.Error(w, "Erro na requisicao", http.StatusInternalServerError)
		return
	}

	//Deletando e testando
	err = controller.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao deletar employee: %v", err)
		http.Error(w, "Erro ao deletar employee", http.StatusBadRequest)
		return
	}

	//Retornando status
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w)
}
