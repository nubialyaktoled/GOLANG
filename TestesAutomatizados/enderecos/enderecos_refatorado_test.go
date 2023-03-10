package enderecos

import "testing"

// comands
// go test -v mostra detalhes
// go test ./... rota tds testes em tds arquivos
// go test --cover mostra cobertura do teste
// go test --coverprofile resultado.txt  -- faz gerar um txt com resultado do teste
// go tool cover --func=resultado.txt -- mostra a cobertura de cada funçao
// go tool cover --html=resultado.txt -- mostra arquivo html que vai mostrar todas linhas cobertas
//HTML output written to /tmp/cover4286831571/coverage.html

type CenariodeTeste struct {
	gotEndereco  string
	wantEndereco string
}

func TestTipoDeEnderecoRefact(t *testing.T) {
	//diz que pode rodar em paralelo
	t.Parallel()

	cenariosDeTeste := []CenariodeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista ABC", "Avenida"},
		{"Rua ABC", "Rua"},
		{"Rua ABC", "Rua"},
		//{"Praça ABC", "Tipo Invalido"},
		{"Rua ABC", "Rua"},
		//{"ABC", "Tipo Invalido"},
		//{"ESTRADA ABC", "Estrada"},
		{"Rua ABC", "Rua"},
		{"Rua ABC", "Rua"},
	}

	for item, cenario := range cenariosDeTeste {
		gotEndereco := TipoDeEndereco(cenario.gotEndereco)
		if gotEndereco != cenario.wantEndereco {
			t.Errorf("O tipo de endereço recebido é diferente do esperado, esperava %s, recebi %s, cenario: %d", cenario.wantEndereco, gotEndereco, item)
		}
	}

}

func TestOutroTeste(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("deu ruim")
	}
}
