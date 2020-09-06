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

// func TestGetOperator(t *testing.T) {
// 	texte := "+="
// 	dataType, longueur := getOperator([]byte(texte), 0)
// 	if dataType != equalAdd {
// 		t.Errorf("Erreur type : " + string(dataType))
// 	}
// 	if longueur != 2 {
// 		t.Errorf("Erreur longueur : " + string(longueur))
// 	}
// }

func TestLexerCondition(t *testing.T) {
	texte := "if (a==5)"
	tokenTab := lexer([]byte(texte))
	if tokenTab[0].dataType != keywordIf {
		t.Errorf("Erreur type : " + string(tokenTab[0].dataType))
	}
	if tokenTab[0].valeurString != "if" {
		t.Errorf("Erreur type : " + tokenTab[0].valeurString)
	}

}
