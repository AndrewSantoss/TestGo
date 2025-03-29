package controller

import (
	database "goApi/database"
	"goApi/model"
)

func FindByCep(cep string) (employees []model.Employee, err error) {
	//Iniciando e testando a conexão com o banco
	conn, err := database.ConnectDB()
	if err != nil {
		return
	}

	//Fazendo o select
	rows, err := conn.Query(`SELECT * FROM EMPLOYEE WHERE CEP=$1`, cep)

	//Transformando a linha em um objeto
	for rows.Next() {
		var employee model.Employee
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Age, &employee.CEP, &employee.Gender, &employee.Street, &employee.District, &employee.City, &employee.State, &employee.Number)
		employees = append(employees, employee)
	}

	//Fechando a conexão
	defer conn.Close()
	return
}
