package bonus

import "fmt"

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	destinos := make(map[string]string)

	for _, caminho := range caminhos {
		origem := caminho[0]
		destino := caminho[1]
		destinos[origem] = destino
	}
	
	for _, destino := range destinos {
		_, ok := destinos[destino]
		if !ok {
			return destino, nil
		}
	}

	return "", fmt.Errorf("nenhum destino final encontrado")
}
