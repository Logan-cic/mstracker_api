package main

import (
	"fmt"
	"log"
	"mstracker_api/src/config"
	"mstracker_api/src/router"
	"net/http"
)

func main() {
	config.Carregar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
