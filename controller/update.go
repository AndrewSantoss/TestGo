package controller

import (
	database "goApi/database"
	"goApi/model"
)

func Update(id int64, employee model.Employee) (err error) {
	//Iniciando e testando a conexão com o banco
	conn, err := database.ConnectDB()
	if err != nil {
		return
	}

	//Fazendo o update
	conn.Exec(`UPDATE EMPLOYEE SET NAME=$1,AGE=$2,CEP=$3,GENDER=$4,STREET=$5,DISTRICT=$6,CITY=$7,STATE=$8,NUMBER=$9 WHERE ID=$10`, employee.Name, employee.Age, employee.CEP, employee.Gender, employee.Street, employee.District, employee.City, employee.State, employee.Number, id)

	//Fechando a conexão
	defer conn.Close()
	return
}
