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
func TestGetOperator(t *testing.T) {
	texte := "3 += 2"
	dataType, longueur := get_operator([]byte(texte), 4)
	if dataType != Equal_add {
		t.Errorf("Erreur type : " + string(dataType))
	}
	if longueur != 2 {
		t.Errorf("Erreur longueur : ")
	}
}
