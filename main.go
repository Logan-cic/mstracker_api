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
	fmt.Println(config.Porta)
	
	fmt.Println(config.StringConexaoBanco)


	fmt.Println("Rodando API...")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}