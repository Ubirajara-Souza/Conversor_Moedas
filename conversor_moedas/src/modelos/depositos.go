package modelos

import (
	"errors"
	"time"
)

// Depositos representa depositos armazenados no conversor de moedas
type Depositos struct {
	ID              uint64    `json:"id,omitempty"`
	ValorDepositado float64   `json:"ValorDepositado,omitempty"`
	DepositadoEm    time.Time `json:"DepositadoEm,omitempty"`
}

// Preparar vai chamar o método para validar o deposito recebido
func (deposito *Depositos) Preparar() error {
	if erro := deposito.validar(); erro != nil {
		return erro
	}

	return nil
}

// validar verificar se o depósito foi feito de forma correta.
func (deposito *Depositos) validar() error {
	if deposito.ValorDepositado <= 0 {
		return errors.New("O valor depósitado é obrigatório e não pode receber valores negativos ou igual a zero")
	}
	return nil
}
