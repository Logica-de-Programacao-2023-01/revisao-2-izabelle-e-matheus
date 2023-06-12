package bonus

import (
	"fmt"
)

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	destinoMap := make(map[string]bool)
	for _, caminho := range caminhos {
		origem := caminho[0]
		destino := caminho[1]
		// Se o destino já foi encontrado anteriormente como origem, não é o destino final
		if destinoMap[destino] {
			destinoMap[destino] = false
		}
		// Marca o destino atual como origem
		destinoMap[origem] = true
	}

	var destinoFinal string
	for destino, isOrigem := range destinoMap {
		if isOrigem {
			// Se o destino atual é uma origem e não foi encontrado como destino, é o destino final
			if destinoFinal != "" {
				// Mais de um destino final encontrado
				return "", fmt.Errorf("múltiplos destinos finais encontrados")
			}
			destinoFinal = destino
		}
	}

	if destinoFinal == "" {
		// Nenhum destino final encontrado
		return "", fmt.Errorf("nenhum destino final encontrado")
	}

	return destinoFinal, nil
}
