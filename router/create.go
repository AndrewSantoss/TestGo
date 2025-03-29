package router

import (
	"encoding/json"
	"goApi/controller"
	"goApi/model"
	"goApi/viacep"
	"log"
	"net/http"
	"regexp"
)

func Create(w http.ResponseWriter, r *http.Request) {
	//Criando uma verialvel
	var employee model.Employee

	//Pegando e testando o body da requisição
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Printf("Erro na requisicao: %v", err)
		http.Error(w, "Erro na requisicao", http.StatusInternalServerError)
		return
	}

	//Verificando itens obrigatorios
	if employee.Name == "" {
		log.Printf("Erro na requisicao, indique o nome")
		http.Error(w, "Erro na requisicao, indique o nome", http.StatusBadRequest)
		return
	}
	if employee.Age == 0 {
		log.Printf("Erro na requisicao, indique a idade")
		http.Error(w, "Erro na requisicao, indique a idade", http.StatusBadRequest)
		return
	}
	if employee.CEP == "" {
		log.Printf("Erro na requisicao, indique o CEP")
		http.Error(w, "Erro na requisicao, indique o CEP", http.StatusBadRequest)
		return
	}

	//Testando se foi informado apenas o CEP
	if employee.CEP != "" {
		if employee.Street == "" || employee.District == "" || employee.City == "" || employee.State == "" {
			//Verificar se o cep tem traço e pegando o endereço
			if regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(employee.CEP) {
				viacep.GetAddress(employee.CEP, &employee)
			} else {
				log.Printf("Erro na requisicao, retire o traço do CEP")
				http.Error(w, "Erro na requisicao, retire o traço do CEP", http.StatusBadRequest)
				return
			}
		}
	}

	//Criando e testando o funcionario
	err = controller.Create(employee)
	if err != nil {
		return
	}

	//Retornando status
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w)
}
