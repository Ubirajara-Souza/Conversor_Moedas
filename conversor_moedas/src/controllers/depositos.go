package controllers

import (
	"conversor_moedas/src/banco"
	"conversor_moedas/src/modelos"
	"conversor_moedas/src/repositorios"
	"conversor_moedas/src/respostas"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FazerDepositos inserir depositos no banco de dados
func FazerDepositos(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var deposito modelos.Depositos
	if erro = json.Unmarshal(corpoRequest, &deposito); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = deposito.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeDepositos(db)

	deposito.ID, erro = repositorio.Inserir(deposito)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, deposito)
	fmt.Println()
	fmt.Printf("Foi depositado o valor de R$ %.2f \n", deposito.ValorDepositado)
}

// ListarDepositos listar todas os depositos salvos no banco
func ListarDepositos(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeDepositos(db)

	depositos, erro := repositorio.ListarDepositos()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, depositos)
}

//BuscarExtratoTotal busca o saldo total e converte todos os depositos salvas no banco
func BuscarExtratoTotal(w http.ResponseWriter, r *http.Request) {
	moeda := r.URL.Query().Get("moeda")

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	repositorio := repositorios.NovoRepositorioDeDepositos(db)

	extratoTotal, erro := repositorio.BuscarExtratoTotal(moeda)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, extratoTotal)
}
