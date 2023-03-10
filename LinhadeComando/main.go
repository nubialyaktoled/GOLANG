package main

import ("fmt"
"log"
"os"
app "GodoZero/app")

func main(){

	//para teste
	//go run main.go ip --host mercadolivre.com.br
	//go run main.go ip 

// Depois de compilar
	//./GodoZero servidores --host mercadolivre.com.br

	fmt.Println("Ponto de partida")
	aplicacao := app.Gerar()
	err := aplicacao.Run(os.Args)
	// os.args = parametro para comandos do sistema operacional seja reconhecido

	if err != nil {
 		 log.Fatal(err)
	}

	

	
}