package main

import (
	"fmt"
	"tipos-validos/enderecos"
)

func main() {

	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
