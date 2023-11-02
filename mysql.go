package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	time.LoadLocation("UTC")
	ConnectionUrl := "golang:golang@/dbgolang?charset=utf8&loc=UTC"
	db, erro := sql.Open("mysql", ConnectionUrl)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Conexão Bem sucedida")

	linhas, erro := db.Query("SELECT * FROM USUARIOS")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}
