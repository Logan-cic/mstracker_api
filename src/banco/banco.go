package banco

import (
	"database/sql"
	"fmt"
	config "mstracker_api/src/configs"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre a conexão com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
    fmt.Println("opaaaaaaaaa")
    // Abra a conexão com o banco de dados MySQL
    db, err := sql.Open("mysql",  config.StringConexaoBanco)

    if err != nil {
        return nil, err
    }

    // Verifique se a conexão com o banco de dados é válida
    if err = db.Ping(); err != nil {
        db.Close()
        return nil, err
    }

    fmt.Println("Passsou do ping uuuhuuu!!!!!!")
    
    return db, nil
}