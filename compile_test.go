package main

import "testing"

func TestOperator(t *testing.T) {
	texte := "3 * 2"
	tokenTab := lexer([]byte(texte))
	if string(tokenTab[0].valeurString) != "3" {
		t.Errorf("Erreur")
	}
	if string(tokenTab[1].valeurString) != "*" {
		t.Errorf("Erreur")
	}

}
