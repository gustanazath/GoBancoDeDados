package main

import (
	"banco-de-dados/BasicCrud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", (server.CriarUsuario)).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", (server.BuscarUsuraios)).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", (server.BuscarUsuraio)).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", (server.AtualizarUsuario)).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", (server.DeleteUsuario)).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
