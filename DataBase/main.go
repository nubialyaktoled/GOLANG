package main

import(
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
func main(){

	urlConexao := "golang:golang@/mydb?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	
	fmt.Println("A conexão esta aberta!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)

}