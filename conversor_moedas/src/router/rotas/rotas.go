package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas do Conversor de Moedas
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasDepositos

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
