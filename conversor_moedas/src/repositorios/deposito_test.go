package repositorios

import (
	"math"
	"testing"
)

func TestCalcularCambio(t *testing.T) {
	ValorEsperado := 12.28
	ValorRecebido := math.Round(100/calcularCambio(6.444)*100) / 100
	if ValorRecebido != ValorEsperado {
		t.Errorf("Função esperava: %f, recebeu: %f", ValorEsperado, ValorRecebido)
	}
}
