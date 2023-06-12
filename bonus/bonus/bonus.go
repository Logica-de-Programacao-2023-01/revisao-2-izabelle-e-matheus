package bonus

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, bool) {
	grafo := make(map[string]string)

	for _, caminho := range caminhos {
		origem := caminho[0]
		destino := caminho[1]
		grafo[origem] = destino
	}

	// Verificar se há um caminho válido
	for _, caminho := range caminhos {
		origem := caminho[0]
		_, ok := grafo[origem]
		if !ok {
			return "", false
		}
	}

	// Encontrar o destino final percorrendo o grafo
	origem := caminhos[0][0]
	for {
		destino, ok := grafo[origem]
		if !ok {
			return "", false
		}

		if destino == origem {
			return "", false
		}

		origem = destino

		_, ok = grafo[origem]
		if !ok {
			break
		}
	}

	return origem, true
}
