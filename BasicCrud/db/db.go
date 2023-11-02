package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar are a conexão com o Banco de dados
func Conectar() (*sql.DB, error) {
	time.LoadLocation("UTC")
	ConnectionUrl := "golang:gola2ng@/dbgolang?charset=utf8&loc=UTC"

	db, erro := sql.Open("mysql", ConnectionUrl)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
