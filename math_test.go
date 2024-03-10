package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(15, 15)

	if total != 30 {
		t.Errorf("resultado da soma e invalido: Resultado %d. Esperado: %d", total, 30)
	}

}

func Testsubtrai(t *testing.T) {

	total := subtrai(15, 15)

	if total != 0 {
		t.Errorf("resultado da subtracao e invalido: Resultado %d. Esperado: %d", total, 0)
	}

}
