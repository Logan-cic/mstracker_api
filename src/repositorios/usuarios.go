package repositorios

import (
	"database/sql"
	"fmt"
	"mstracker_api/src/modelos"
)

// usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}


// NovoRepositorioDeUsuario cria um repositorio de usuarios
func NovoRepositorioDeUsuario(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}



// Cria insere um usuario no banco de dados
func (repositorio Usuarios) Cria(usuario modelos.Usuario) (uint64, error) {
	
	fmt.Println("vou inserir!!!!")

	statement, err := repositorio.db.Prepare(
		"insert into usuarios (nome, email, senha) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Email, usuario.Senha)
	
	if err != nil {
		return 0, err
	}
	
	ultimoIDInserido, err := resultado.LastInsertId()

	if err != nil {
		return 0, err
	}
	

	return uint64(ultimoIDInserido), nil
}