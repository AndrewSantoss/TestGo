Sobre: 
Criação de um sistema de gerenciamento de funcionarios de uma empresa pequena que permita cadastrar, buscar, atualizar e deletar um funcionário.

Caracteristicas: 
Linguagem Go - Versão 1.24.1

Dependencias utilizadas: 
1 - chi
2 - pq

Informações: 
/API - Pasta que contem a main
/controller - Pasta que contem as chamadas do banco de dados
/database - Pasta que contem a função de conexão com o banco de dados
/model - Pasta que contem a criação dos objetos
/router - Pasta que contem as rodas que cada endpoint vai passar
/viacep - Pasta que contem a integração com o viacep
docker-compose.yml - Inicia o banco
Dockerfile.yml - Fiquei com muitas duvidas sobre o funcionamento deste processo e como não era um item obrigatorio acabei pulando kkkkk
EMPLOYEES GO.postman_collection.json - Referente a API utilizada nos testes
go.mod e go.sum - Arquivos que contem as dependencias utilizadas
init.sql - Arquivo que utilizo no docker-compose.yml para subir o banco

Itens:
1ª Etapa: Ok
2ª Etapa: Ok
3ª Etapa: Ok
4ª Etapa: Ok
5ª Etapa: Ok
6ª Etapa: Ok

Iniciando a aplicação:
1 - Verifique se você está com a versão 1.24.1 do GO
2 - Rode o arquivo docker-compose.yml ou no terminal rode o comando "docker-compose up"
3 - Se tudo estiver certo a seguinte mensagem vai aparecer no terminal "------------------ Conexao Concluida ---------------"

Observação:
Caso o dockerfile ou o compose não funcionar essas são as informações...

Server name: go_db
Host name: localhost
Port: 5432
Databese: postgres
User: adm
Pass: 123

CREATE TABLE EMPLOYEE(
	ID SERIAL PRIMARY KEY,
	NAME VARCHAR(100) NOT NULL,
	AGE INT NOT NULL,
	CEP VARCHAR(10) NOT NULL,
	GENDER VARCHAR(10),
	STREET VARCHAR(50),
	DISTRICT VARCHAR(50),
	CITY VARCHAR(50),
	STATE VARCHAR(50),
	NUMBER INT
);

INSERT INTO EMPLOYEE(NAME,AGE,CEP,GENDER,STREET,DISTRICT,CITY,STATE,NUMBER) VALUES ('Andrew Alves dos Santos',25,'06655110','Masculino','Rua Andradina','Jardim Santo Américo','Itapevi','SP',123)
INSERT INTO EMPLOYEE(NAME,AGE,CEP,GENDER,STREET,DISTRICT,CITY,STATE,NUMBER) VALUES ('Julia Sabrina Josefa Gonçalves',59,'07190913','Feminino','Rua Célia Domingues Faustino','Parque Cecap','Guarulhos','SP',330) 
INSERT INTO EMPLOYEE(NAME,AGE,CEP,GENDER,STREET,DISTRICT,CITY,STATE,NUMBER) VALUES ('Geraldo Julio Carlos Alves',35,'49061052','Masculino','Rua Esplanada','Santo Antônio','Aracaju','SE',104) 
INSERT INTO EMPLOYEE(NAME,AGE,CEP,GENDER,STREET,DISTRICT,CITY,STATE,NUMBER) VALUES ('Regina Sebastiana Allana Moraes',20,'69915120','Feminino','Travessa No Limite II','Conjunto Esperança','Rio Branco','AC',378)


