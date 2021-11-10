package main

import (
	"conversor_moedas/src/config"
	"conversor_moedas/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Rodando Conversor de Moeda na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
