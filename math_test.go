package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(15, 15)
	valorEsperado := 10

	if total != valorEsperado {
		t.Errorf("Resultado da soma é inválido! Valor esperado é %d e o resultado foi %d.", valorEsperado, total)
	}

}
