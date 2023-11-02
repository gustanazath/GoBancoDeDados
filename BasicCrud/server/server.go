package server

import (
	"banco-de-dados/BasicCrud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:id`
	Nome  string `json:nome`
	Email string `json:email`
}

// CRIA USUARIO NA BASE DE DADOS
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	RequestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler a request"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(RequestBody, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao converter conectar ao banco"))
		return
	}

	statement, erro := db.Prepare("insert into usuarios(nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ai criar o statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter Id inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idInserido)))
}

// BUSCAR TODOS OS USUARIOS NA BASE DE DADOS
func BuscarUsuraios(w http.ResponseWriter, r *http.Request) {
	db, erro := db.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com a base"))
	}
	defer db.Close()

	lines, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuario"))
	}
	defer lines.Close()

	var usuarios []usuario

	for lines.Next() {
		var usuario usuario
		if erro := lines.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuarios para JSON"))
		return
	}
}

// BUSCAR USUARIOS ESPECIFICO NA BASE DE DADOS
func BuscarUsuraio(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	ID, erro := strconv.ParseUint(parameters["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	line, erro := db.Query("SELECT * FROM usuarios Where id = ?", ID)

	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}
	var usuario usuario
	if line.Next() {
		if erro := line.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear usuário"))
			return
		}
		if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
			w.Write([]byte("Erro ao converter o usuário para JSON!"))
			return
		}
	}
}

// ATUALIZA USUÁRIO NO BANCO DE DADOS
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, erro := strconv.ParseUint(parameters["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro"))
		return
	}

	BodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o body"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(BodyRequest, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar na base"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuario"))
	}

	w.WriteHeader(http.StatusNoContent)
}

// DELETAR UM USUARIO NO BANCO DE DADOS
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)
	ID, erro := strconv.ParseUint(parameters["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro para inteiro"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectaro a base"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM USUARIOS WHERE id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
