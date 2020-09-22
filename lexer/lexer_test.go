package lexer

import (
	"testing"

	token "github.com/StevenK8/Compilateur/token"
)

func TestOperator(t *testing.T) {
	texte := "3 * 2"
	tokenTab := Lexer([]byte(texte))
	if string(tokenTab[0].ValeurString) != "3" {
		t.Errorf("Erreur")
	}
	if string(tokenTab[1].ValeurString) != "*" {
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
	tokenTab := Lexer([]byte(texte))
	if tokenTab[0].DataType != token.KeywordIf {
		t.Errorf("Erreur type : " + string(tokenTab[0].DataType))
	}
	if tokenTab[0].ValeurString != "if" {
		t.Errorf("Erreur type : " + tokenTab[0].ValeurString)
	}

}
