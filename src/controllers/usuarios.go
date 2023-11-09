package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mstracker_api/src/banco"
	"mstracker_api/src/modelos"
	"mstracker_api/src/repositorios"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//CriarUsuario insere um usuário no banco de dados
func CriarUsuario( w http.ResponseWriter, r *http.Request){
	corpoRequest, err := io.ReadAll(r.Body)

	if err != nil {
		
		log.Fatal()

	}
	var usuario modelos.Usuario

	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Chamou a conexão com o banco")
	// sql.Open("mysql", "root:mstracker@tcp(172.17.0.2:3306)/users_db")

	db, err := banco.Conectar()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectou!")

	repositorio := repositorios.NovoRepositorioDeUsuario(db)

	fmt.Println("POOOOOOORAAAAAAAAAAA")

	repositorio.Cria(usuario)
	if err != nil {
		log.Fatal()
	}
	w.Write([]byte(fmt.Sprintf("Usuario inserido")))
}
//BuscarUsuario procura por todos os usuário no banco de dados
func BuscarUsuarios( w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando todos os usuários"))
}

//BuscarUsuario procura por um usuário no banco de dados
func BuscarUsuario( w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando um Usuário"))
}

//AtulizaUsuario atualiza informações sobre um usuário no banco de dados
func AtualizarUsuario( w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando Usuário"))
}

//DeletarUsuario deleta um usuário no banco de dados
func DeletarUsuario( w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletando Usuário"))
}