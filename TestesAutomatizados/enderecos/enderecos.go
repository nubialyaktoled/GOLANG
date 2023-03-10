package enderecos

import "strings"

// TipoDeEndereco verifica se endereço tem endereço valido
func TipoDeEndereco(endereco string) string {

	tiposValidos := []string{"rua", "avenida", "estradas", "rodovias"}

	//ToLower = deixa pequeno
	enderecoLetraMinuscula := strings.ToLower(endereco)

	//função split divide string em um slice sempre que tiver um espaço
	//ex: Rua da Graça = ["Rua", "da", "Graça"]
	primeiraPalavraDoEndereco := strings.Split(enderecoLetraMinuscula, " ")[0]

	enderecoValido := false
	for _, item := range tiposValidos {
		if item == primeiraPalavraDoEndereco {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Invalido"

}
