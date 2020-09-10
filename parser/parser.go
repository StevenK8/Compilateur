package parser

import (
	"fmt"
	"log"

	token "Compilateur/token"
)

type typeNoeud string

const (
	NoeudExpo   typeNoeud = "NoeudExpo"
	NoeudMult   typeNoeud = "NoeudMult"
	NoeudDiv    typeNoeud = "NoeudDiv"
	NoeudMod    typeNoeud = "NoeudMod"
	NoeudAdd    typeNoeud = "NoeudAdd"
	NoeudSub    typeNoeud = "NoeudSub"
	NoeudDec    typeNoeud = "NoeudDeclaration"
	NoeudRef    typeNoeud = "NoeudReference"
	NoeudAffect typeNoeud = "NoeudAffection"

	NoeudLessThan    typeNoeud = "NoeudLessThan"
	NoeudGreaterThan typeNoeud = "NoeudGreaterThan"
	NoeudEqualequal  typeNoeud = "NoeudEqualequal"
	NoeudNotEqual    typeNoeud = "NoeudNotEqual"

	NoeudSubUnaire typeNoeud = "NoeudSubUnaire"
	NoeudNot       typeNoeud = "NoeudNot"

	NoeudAnd typeNoeud = "NoeudAnd"
	NoeudOr  typeNoeud = "NoeudOr"

	NoeudOpenParen  typeNoeud = "NoeudOpen_Paren"
	NoeudCloseParen typeNoeud = "NoeudClose_Paren"
	NoeudConst      typeNoeud = "NoeudConst"

	NoeudDebug typeNoeud = "NoeudDebug"

	NoeudBlock typeNoeud = "NoeudBlock"
	NoeudDrop  typeNoeud = "NoeudDrop"
	NoeudTest  typeNoeud = "NoeudTest"
)

type operation struct {
	typeOfToken token.TokenType
	prioD       int
	prioG       int
	typeOfNoeud typeNoeud
}

var tabOperation = []operation{
	//	operation {token.Exposant, 60, 61, NoeudExpo},
	operation{token.Not, 55, 56, NoeudNot},
	operation{token.MinusUnaire, 55, 56, NoeudSubUnaire},
	operation{token.OperatorMult, 50, 51, NoeudMult},
	operation{token.OperatorDiv, 50, 51, NoeudDiv},
	//	operation {token.Modulo, 50, 51, NoeudMod},
	operation{token.OperatorPlus, 40, 41, NoeudAdd},
	operation{token.OperatorMinus, 40, 41, NoeudSub},
	operation{token.LessThan, 30, 31, NoeudLessThan},
	operation{token.GreaterThan, 30, 31, NoeudGreaterThan},
	operation{token.Equalequal, 30, 31, NoeudEqualequal},
	operation{token.NotEqual, 30, 31, NoeudNotEqual},
	operation{token.And, 20, 21, NoeudAnd},
	operation{token.Or, 10, 11, NoeudOr},
	operation{token.Equal, 5, 5, NoeudAffect},
}

func estPrio(typeWanted token.TokenType) bool {
	for _, prioBase := range tabOperation {
		if typeWanted == prioBase.typeOfToken {
			return true
		}
	}
	return false
}

func getPrio(typeWanted token.TokenType) operation {
	for _, prioBase := range tabOperation {
		if typeWanted == prioBase.typeOfToken {
			return prioBase
		}
	}

	log.Fatal("[ Line ", courant().NbLigne, " ] Erreur : operateur ", typeWanted, " non trouvé.")
	return operation{"", 0, 0, ""}
}

type Noeud struct {
	TypeDeNoeud   typeNoeud
	NbLigne       int
	Fils          []Noeud
	ValeurEntiere int
	ValeurString  string
	Slot          int
}

var tokenTab []token.Token
var posToken int

// Parser : Main function of this package
func Parser(afterLexer []token.Token) Noeud {
	tokenTab = afterLexer
	var mainNoeud Noeud = instruction()

	//PrintNoeud(mainNoeud, 0)
	return mainNoeud
}

func PrintNoeud(N Noeud, decalage int) {
	var decal string
	for i := 0; i < decalage; i++ {
		decal += "\t"
	}
	fmt.Println(decal, "\\_", "Noeud : ", N.TypeDeNoeud, " - ", N.ValeurEntiere, " - Slot : ", N.Slot)
	for _, child := range N.Fils {
		PrintNoeud(child, decalage+1)
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

func ajouterEnfant(parent Noeud, childs ...Noeud) Noeud {
	for _, child := range childs {
		parent.Fils = append(parent.Fils, child)
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

func nouveauNoeud(typeDuNoeud typeNoeud, LineNumber int) Noeud {
	return Noeud{typeDuNoeud, LineNumber, nil, 0, "", 0}
}

func atome() Noeud {
	var N Noeud
	if verifier(token.ParentheseOuvrante) {
		N = expression(0)
		accepter(token.ParentheseFermante)
		return N
	} else if verifier(token.OperatorMinus) {
		N = nouveauNoeud(NoeudSubUnaire, courant().NbLigne)
		A := expression(getPrio(token.MinusUnaire).prioD)
		N = ajouterEnfant(N, A)
		return N
	} else if verifier(token.Not) {
		N = nouveauNoeud(NoeudNot, courant().NbLigne)
		A := expression(getPrio(token.Not).prioD)
		N = ajouterEnfant(N, A)
		return N
	} else if courant().DataType == token.Constant {
		N = nouveauNoeud(NoeudConst, courant().NbLigne)
		N.ValeurEntiere = courant().ValeurInt
		avancer()
		return N
	} else if courant().DataType == token.Ident {
		N = nouveauNoeud(NoeudRef, courant().NbLigne)
		N.ValeurString = courant().ValeurString
		avancer()
		return N
	}

	log.Fatal("[ Line ", courant().NbLigne, " ] Erreur : atome - ", token.Constant, " attendu, ", courant().DataType, " reçu.")
	return N
}

func expression(prioMin int) Noeud {
	var N Noeud = atome()
	if estPrio(courant().DataType) {
		for getPrio(courant().DataType).prioG > prioMin {
			var op operation = getPrio(courant().DataType)
			var line int = courant().NbLigne
			avancer()
			var A Noeud = expression(op.prioD)
			var T Noeud = nouveauNoeud(op.typeOfNoeud, line)
			T = ajouterEnfant(T, N, A)
			N = T
			if !estPrio(courant().DataType) {
				break
			}
		}
	}
	return N
}

func instruction() Noeud {
	var N Noeud

	if verifier(token.Debug) {
		E := expression(0)
		accepter(token.PointVirgule)
		N = nouveauNoeud(NoeudDebug, courant().NbLigne)
		N = ajouterEnfant(N, E)
		return N

	} else if verifier(token.LeftBrace) {
		N = nouveauNoeud(NoeudBlock, courant().NbLigne)
		for !verifier(token.RightBrace) {
			N = ajouterEnfant(N, instruction())
		}
		return N

	} else if verifier(token.KeywordIf) {
		accepter(token.ParentheseOuvrante)
		E1 := expression(0)
		accepter(token.ParentheseFermante)
		I1 := instruction()
		N = nouveauNoeud(NoeudTest, courant().NbLigne)
		N = ajouterEnfant(N, E1, I1)
		if verifier(token.KeywordElse) {
			I2 := instruction()
			N = ajouterEnfant(N, I2)
		}
		return N
	} else if verifier(token.KeywordInt) {
		N = nouveauNoeud(NoeudDec, courant().NbLigne)
		N.ValeurString = courant().ValeurString
		avancer()
		accepter(token.PointVirgule)
		return N
	} else {
		E1 := expression(0)
		accepter(token.PointVirgule)
		N = nouveauNoeud(NoeudDrop, courant().NbLigne)
		N = ajouterEnfant(N, E1)
		return N
	}

}
