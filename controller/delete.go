package controller

import (
	database "goApi/database"
)

func Delete(id int64) (err error) {
	//Iniciando e testando a conexão com o banco
	conn, err := database.ConnectDB()
	if err != nil {
		return
	}

	//Fazendo o delete
	conn.Exec(`DELETE FROM EMPLOYEE WHERE ID = $1`, id)

	//Fechando a conexão
	defer conn.Close()
	return
}
