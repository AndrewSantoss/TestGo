package viacep

import (
	"encoding/json"
	"fmt"
	"goApi/model"
	"net/http"
)

func GetAddress(cep string, employee *model.Employee) (err error) {
	//Pegando e testando o endereço
	strGet := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", cep)
	resp, err := http.Get(strGet)
	if err != nil {
		return
	}

	//Transformando a resposta em um objeto
	var address model.Viacep
	err = json.NewDecoder(resp.Body).Decode(&address)
	if err != nil {
		return
	}

	//Atualizando o endereço
	employee.Street = address.Logradouro
	employee.District = address.Bairro
	employee.City = address.Localidade
	employee.State = address.UF

	return

}
