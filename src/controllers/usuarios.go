package controllers

import "net/http"


//CriarUsuario insere um usuário no banco de dados
func CriarUsuario( w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Criando Usuário"))
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