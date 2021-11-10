package rotas

import (
	"conversor_moedas/src/controllers"
	"net/http"
)

var rotasDepositos = []Rota{
	{
		URI:    "/depositar",
		Metodo: http.MethodPost,
		Funcao: controllers.FazerDepositos,
	},

	{
		URI:    "/depositos",
		Metodo: http.MethodGet,
		Funcao: controllers.ListarDepositos,
	},

	{
		URI:    "/saldo",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarExtratoTotal,
	},
}
