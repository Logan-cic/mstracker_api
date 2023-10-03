package main

import (
	"fmt"
	"log"
	"mstracker_api/src/router"
	"net/http"
)

func main() {
	fmt.Println("Rodando API...")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}