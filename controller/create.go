package controller

import (
	database "goApi/database"
	"goApi/model"
)

func Create(employee model.Employee) (err error) {
	//Iniciando e testando a conexão com o banco
	conn, err := database.ConnectDB()
	if err != nil {
		return
	}

	//Fazendo o insert
	conn.QueryRow(`INSERT INTO EMPLOYEE(NAME,AGE,CEP,GENDER,STREET,DISTRICT,CITY,STATE,NUMBER) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`, employee.Name, employee.Age, employee.CEP, employee.Gender, employee.Street, employee.District, employee.City, employee.State, employee.Number)

	//Fechando a conexão
	defer conn.Close()
	return
}
