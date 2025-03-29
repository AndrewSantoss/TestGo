package controller

import (
	"fmt"
	database "goApi/database"
)

func HalfUpdate(id int64, attribute string, value string) (err error) {
	//Iniciando e testando a conexão com o banco
	conn, err := database.ConnectDB()
	if err != nil {
		return
	}

	//Criando a string e fazendo o update
	str := fmt.Sprintf("UPDATE EMPLOYEE SET %v='%v' WHERE ID=%v", attribute, value, id)
	conn.Exec(str)

	//Fechando a conexão
	defer conn.Close()
	return
}
