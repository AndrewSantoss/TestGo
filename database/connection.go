package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	//String de conexão
	connStr := "user=adm password=123 dbname=postgres host=localhost sslmode=disable"

	//Abrir a conexão
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Erro ao abrir a conexão com o banco", err)
	}

	return conn, err
}
