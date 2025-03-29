package router

import (
	"encoding/json"
	"goApi/controller"
	"goApi/model"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HalfUpdate(w http.ResponseWriter, r *http.Request) {
	//Criando uma variavel
	var employee model.Employee

	//Pegando e testando o id passado na requisição
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro na requisicao: %v", err)
		http.Error(w, "Erro na requisicao", http.StatusInternalServerError)
		return
	}

	//Pegando e testando o body da requisição
	err = json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Printf("Erro na requisicao: %v", err)
		http.Error(w, "Erro na requisicao", http.StatusInternalServerError)
		return
	}

	//Verificando qual atributo tem conteudo e fazendo atualização
	if employee.Name != "" {
		controller.HalfUpdate(int64(id), "NAME", employee.Name)
	}
	if employee.Age != 0 {
		controller.HalfUpdate(int64(id), "AGE", strconv.Itoa(employee.Age))
	}
	if employee.CEP != "" {
		if regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(employee.CEP) {
			controller.HalfUpdate(int64(id), "CEP", employee.CEP)
		} else {
			log.Printf("Erro na requisicao, retire o traço do CEP")
			http.Error(w, "Erro na requisicao, retire o traço do CEP", http.StatusBadRequest)
			return
		}
	}
	if employee.Gender != "" {
		controller.HalfUpdate(int64(id), "GENDER", employee.Gender)
	}
	if employee.Street != "" {
		controller.HalfUpdate(int64(id), "STREET", employee.Street)
	}
	if employee.District != "" {
		controller.HalfUpdate(int64(id), "DISTRICT", employee.District)
	}
	if employee.City != "" {
		controller.HalfUpdate(int64(id), "CITY", employee.City)
	}
	if employee.State != "" {
		controller.HalfUpdate(int64(id), "STATE", employee.State)
	}
	if employee.Number != 0 {
		controller.HalfUpdate(int64(id), "NUMBER", strconv.Itoa(employee.Number))
	}

	return
}
