package repositorios

import (
	"conversor_moedas/src/modelos"
	"database/sql"
	"log"
	"math"
)

// Depositos representa um repositório de depositos
type Depositos struct {
	db *sql.DB
}

// NovoRepositorioDeDepositos cria um repositório de depositos
func NovoRepositorioDeDepositos(db *sql.DB) *Depositos {
	return &Depositos{db}
}

// Inserir insere depositos no banco de dados
func (repositorio Depositos) Inserir(deposito modelos.Depositos) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into depositos (valorDepositado) values(?)",
	)
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()

	resultado, erro := statement.Exec(deposito.ValorDepositado)
	if erro != nil {
		return 0, nil
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	return uint64(ultimoIDInserido), nil
}

//ListarDepositos traz todos os depósitos realizados.
func (repositorio Depositos) ListarDepositos() ([]modelos.Depositos, error) {
	linhas, erro := repositorio.db.Query("select id, valorDepositado, depositadoEm from depositos")
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var depositos []modelos.Depositos
	for linhas.Next() {
		var deposito modelos.Depositos
		if erro = linhas.Scan(
			&deposito.ID,
			&deposito.ValorDepositado,
			&deposito.DepositadoEm,
		); erro != nil {
			return nil, erro
		}
		depositos = append(depositos, deposito)
	}
	return depositos, nil
}

//BuscarExtratoTotal soma os depositos realizados e retorna o saldo total.
func (repositorio Depositos) BuscarExtratoTotal(moeda string) (modelos.Extratos, error) {
	linhas, erro := repositorio.db.Query("select id, valorDepositado, depositadoEm from depositos")
	if erro != nil {
		return modelos.Extratos{}, erro
	}
	defer linhas.Close()

	depositosEmSlice := make([]float64, 3)
	for linhas.Next() {
		var deposito modelos.Depositos
		if erro = linhas.Scan(
			&deposito.ID,
			&deposito.ValorDepositado,
			&deposito.DepositadoEm,
		); erro != nil {
			return modelos.Extratos{}, erro
		}
		depositosEmSlice = append(depositosEmSlice, deposito.ValorDepositado)
	}
	var valorTotalEmBRL float64
	for index, _ := range depositosEmSlice {
		valorTotalEmBRL += depositosEmSlice[index]
	}
	var saldo modelos.Extratos

	if moeda == "EUR" {
		saldo.SaldoMoedaEUR = math.Round(valorTotalEmBRL/calcularCambio(6.444)*100) / 100
	}
	if moeda == "GBP" {
		saldo.SaldoMoedaGBP = math.Round(valorTotalEmBRL/calcularCambio(7.638)*100) / 100
	}
	if moeda == "USD" {
		saldo.SaldoMoedaUSD = math.Round(valorTotalEmBRL/calcularCambio(5.552)*100) / 100
	}
	if moeda == "" {
		saldo.SaldoTotal = math.Round(valorTotalEmBRL*100) / 100
	}
	return saldo, nil
}

//calcularCambio faz as operações de calcular Spread, IOF e Taxa de Câmbio e retorna o valor real da moeda com impostos.
func calcularCambio(moeda float64) float64 {
	if moeda <= 0 {
		log.Fatal("o valor da moeda não pode ser menor ou igual a zero")
	}

	txC := 0.16
	iof := 0.0638
	spread := 0.04

	valorMoeda := moeda + (moeda * spread) + (moeda * iof) + (moeda * txC)

	return valorMoeda
}
