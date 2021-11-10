package modelos

// Extratos representa os saldos dos depositos armazenados no conversor de moedas
type Extratos struct {
	SaldoTotal    float64 `json:"ExtratoTotal,omitempty"`
	SaldoMoedaUSD float64 `json:"ExtratoTotalEmDolar,omitempty"`
	SaldoMoedaEUR float64 `json:"ExtratoTotalEmEuro,omitempty"`
	SaldoMoedaGBP float64 `json:"ExtratoTotalEmLibra,omitempty"`
}
