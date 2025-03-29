package controller

import (
	database "goApi/database"
	"goApi/model"
)

func FindOne(id int64) (employee model.Employee, err error) {
	//Iniciando e testando a conexão com o banco
	conn, err := database.ConnectDB()
	if err != nil {
		return
	}

	//Fazendo o select
	row := conn.QueryRow(`SELECT * FROM EMPLOYEE WHERE ID = $1`, id)

	//Transformando a linha em um objeto
	err = row.Scan(&employee.ID, &employee.Name, &employee.Age, &employee.CEP, &employee.Gender, &employee.Street, &employee.District, &employee.City, &employee.State, &employee.Number)

	//Fechando a conexão
	defer conn.Close()
	return
}
