package main

import (
	"fmt"
	database "goApi/database"
	"goApi/router"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//Iniciando e testando a conexão com o banco
	conn, err := database.ConnectDB()
	err = conn.Ping()
	if err != nil {
		panic(fmt.Sprint("Erro de conexao: %1", err))
	} else {
		fmt.Println("------------------ Conexao Concluida ---------------")
	}

	//Criando as rotas dos endpoint
	r := chi.NewRouter()
	r.Get("/employee/{id}", router.FindOne)                       //ok
	r.Get("/employee", router.FindAll)                            //ok
	r.Get("/employee/cep/{cep}", router.FindByCep)                //ok
	r.Delete("/employee/operation/delete/{id}", router.Delete)    //ok
	r.Post("/employee/operation/create", router.Create)           //ok
	r.Put("/employee/operation/update/{id}", router.Update)       //ok
	r.Patch("/employee/operation/update/{id}", router.HalfUpdate) //ok

	//Indicando o ip/porta para o servidor escultar
	http.ListenAndServe("127.0.0.1:8080", r)
}
