package parser

import (
	"fmt"
	"log"

	token "../token"
)

type typeNoeud string

const (
	Noeud_Expo typeNoeud = "Noeud_Expo"
	Noeud_Mult typeNoeud = "Noeud_Mult"
	Noeud_Div  typeNoeud = "Noeud_Div"
	Noeud_Mod  typeNoeud = "Noeud_Mod"
	Noeud_Add  typeNoeud = "Noeud_Add"
	Noeud_Sub  typeNoeud = "Noeud_Sub"

	Noeud_LessThan    typeNoeud = "Noeud_LessThan"
	Noeud_GreaterThan typeNoeud = "Noeud_GreaterThan"
	Noeud_Equalequal  typeNoeud = "Noeud_Equalequal"
	Noeud_NotEqual    typeNoeud = "Noeud_NotEqual"

	Noeud_And typeNoeud = "Noeud_And"
	Noeud_Or  typeNoeud = "Noeud_Or"

	Noeud_Open_Paren  typeNoeud = "Noeud_Open_Paren"
	Noeud_Close_Paren typeNoeud = "Noeud_Close_Paren"
	Noeud_Const       typeNoeud = "Noeud_Const"
)

type operation struct {
	typeOfToken token.TokenType
	prioD       int
	prioG       int
	typeOfNoeud typeNoeud
}

var tab_operation = []operation{
	//	operation {token.Exposant, 60, 61, Noeud_Expo},
	operation{token.OperatorMult, 50, 51, Noeud_Mult},
	operation{token.OperatorDiv, 50, 51, Noeud_Div},
	//	operation {token.Modulo, 50, 51, Noeud_Mod},
	operation{token.OperatorPlus, 40, 41, Noeud_Add},
	operation{token.OperatorMinus, 40, 41, Noeud_Sub},
	operation{token.LessThan, 30, 31, Noeud_LessThan},
	operation{token.GreaterThan, 30, 31, Noeud_GreaterThan},
	operation{token.Equalequal, 30, 31, Noeud_Equalequal},
	operation{token.NotEqual, 30, 31, Noeud_NotEqual},
	// operation {token.And, 20, 21, Noeud_And},
	// operation {token.Or, 10, 11, Noeud_Or}
}

func estPrio(typeWanted token.TokenType) bool {
	for _, prioBase := range tab_operation {
		if typeWanted == prioBase.typeOfToken {
			return true
		}
	}
	return false
}

func getPrio(typeWanted token.TokenType) operation {
	for _, prioBase := range tab_operation {
		if typeWanted == prioBase.typeOfToken {
			return prioBase
		}
	}

	log.Fatal("[ Line ", courant().NbLigne, " ] Erreur : operateur ", typeWanted, " non trouvé.")
	return operation{"", 0, 0, ""}
}

type noeud struct {
	typeDeNoeud   typeNoeud
	nbLigne       int
	fils          []noeud
	valeurEntiere int
}

var tokenTab []token.Token
var posToken int

// Parser : Main function of this package
func Parser(afterLexer []token.Token) {
	tokenTab = afterLexer
	var mainNoeud noeud = expression(0)

	printNoeud(mainNoeud, 0)
}

func printNoeud(N noeud, decalage int) {
	var decal string = ""
	for i := 0; i < decalage; i++ {
		decal += "\t"
	}
	fmt.Println(decal, "Noeud : ", N.typeDeNoeud, " - ", N.valeurEntiere)
	for _, child := range N.fils {
		printNoeud(child, decalage+1)
	}
}

func avancer() {
	posToken++
}

func courant() token.Token {
	if posToken < len(tokenTab) {
		return tokenTab[posToken]
	}
	return token.Token{token.EOF, "", 0, -1}
}

func ajouterEnfant(parent noeud, childs ...noeud) noeud {
	for _, child := range childs {
		parent.fils = append(parent.fils, child)
	}
	return parent
}

func verifier(typeCheck token.TokenType) bool {
	if courant().DataType == typeCheck {
		avancer()
		return true
	}
	return false
}

func accepter(typeCheck token.TokenType) {
	if courant().DataType != typeCheck {
		log.Fatal("[ Line ", courant().NbLigne, " ] Erreur accepter : type - ", typeCheck, " attendu, ", courant().DataType, " reçu.")
	}
	avancer()
}

func nouveauNoeud(typeDuNoeud typeNoeud, LineNumber int) noeud {
	return noeud{typeDuNoeud, LineNumber, nil, 0}
}

func atome() noeud {
	var N noeud
	if verifier(token.ParentheseOuvrante) {
		N = expression(0)
		accepter(token.ParentheseFermante)
		return N
	}

	if courant().DataType == token.Constant {
		N = nouveauNoeud(Noeud_Const, courant().NbLigne)
		N.valeurEntiere = courant().ValeurInt
		avancer()
		return N
	}

	log.Fatal("[ Line ", courant().NbLigne, " ] Erreur : atome - ", token.Constant, " attendu, ", courant().DataType, " reçu.")
	return N
}

func expression(prioMin int) noeud {
	var N noeud = atome()
	if estPrio(courant().DataType) {
		for getPrio(courant().DataType).prioG > prioMin {
			var op operation = getPrio(courant().DataType)
			var line int = courant().NbLigne
			avancer()
			var A noeud = expression(op.prioD)
			var T noeud = nouveauNoeud(op.typeOfNoeud, line)
			T = ajouterEnfant(T, N, A)
			N = T
			if !estPrio(courant().DataType) {
				break
			}
		}
	}
	return N
}
