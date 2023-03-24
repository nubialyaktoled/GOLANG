package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o Mysql
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	urlConexao := "golang:golang@/mydb?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
