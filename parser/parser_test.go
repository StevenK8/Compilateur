package parser

import (
	lexer "Compilateur/lexer"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	texte := "3 * 2;"
	tokenTab := lexer.Lexer([]byte(texte))

	p := Parser(tokenTab)
	if p.TypeDeNoeud != NoeudMult {
		fmt.Println("Erreur mult")
	}
}
