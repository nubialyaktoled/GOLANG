package enderecos

// TESTE DE UNIDADE, TESTA PEQUENA PARTE DO CÓDIGO

import (
	"testing"
)

// tem que começar com "Test" e ter letra maiuscula na segunda palavra
// parametro é um ponteiro do tipo t, pacote padrão para desenvolver testes
func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Rua Paulista"
	wantEndereco := "Rua"

	gotEndereco := TipoDeEndereco(enderecoParaTeste)

	if wantEndereco != gotEndereco {
		t.Error("O tipo de endereço recebido é diferente do esperado")
	}

}
