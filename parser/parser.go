package parser

import (
	"fmt"
	"log"

	token "github.com/StevenK8/Compilateur/token"
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

	NoeudLessThan       typeNoeud = "NoeudLessThan"
	NoeudGreaterThan    typeNoeud = "NoeudGreaterThan"
	NoeudEqualequal     typeNoeud = "NoeudEqualequal"
	NoeudNotEqual       typeNoeud = "NoeudNotEqual"
	NoeudLessOrEqual    typeNoeud = "NoeudLessOrEqual"
	NoeudGreaterOrEqual typeNoeud = "NoeudGreaterOrEqual"

	NoeudSubUnaire typeNoeud = "NoeudSubUnaire"
	NoeudNot       typeNoeud = "NoeudNot"

	NoeudAnd typeNoeud = "NoeudAnd"
	NoeudOr  typeNoeud = "NoeudOr"

	NoeudOpenParen  typeNoeud = "NoeudOpen_Paren"
	NoeudCloseParen typeNoeud = "NoeudClose_Paren"
	NoeudConst      typeNoeud = "NoeudConst"

	NoeudDebug  typeNoeud = "NoeudDebug"
	NoeudReturn typeNoeud = "NoeudReturn"
	NoeudSend   typeNoeud = "NoeudSend"

	NoeudBlock typeNoeud = "NoeudBlock"
	NoeudDrop  typeNoeud = "NoeudDrop"
	NoeudTest  typeNoeud = "NoeudTest"

	NoeudLoop     typeNoeud = "NoeudLoop"
	NoeudBreak    typeNoeud = "NoeudBreak"
	NoeudContinue typeNoeud = "NoeudContinue"

	NoeudAppel       typeNoeud = "NoeudAppel"
	NoeudFonction    typeNoeud = "NoeudFonction"
	NoeudIndirection typeNoeud = "NoeudIndirection"
)

type operation struct {
	typeOfToken token.TokenType
	prioD       int
	prioG       int
	typeOfNoeud typeNoeud
}

var tabOperation = []operation{
	operation{token.Pow, 60, 61, NoeudExpo},
	operation{token.Not, 55, 56, NoeudNot},
	operation{token.Pointeur, 55, 56, NoeudIndirection},
	operation{token.MinusUnaire, 55, 56, NoeudSubUnaire},
	operation{token.OperatorMult, 50, 51, NoeudMult},
	operation{token.OperatorDiv, 50, 51, NoeudDiv},
	operation{token.OperatorMod, 50, 51, NoeudMod},
	operation{token.OperatorPlus, 40, 41, NoeudAdd},
	operation{token.OperatorMinus, 40, 41, NoeudSub},
	operation{token.LessThan, 30, 31, NoeudLessThan},
	operation{token.GreaterThan, 30, 31, NoeudGreaterThan},
	operation{token.Equalequal, 30, 31, NoeudEqualequal},
	operation{token.NotEqual, 30, 31, NoeudNotEqual},
	operation{token.LessOrEqual, 30, 31, NoeudLessOrEqual},
	operation{token.GreaterOrEqual, 30, 31, NoeudGreaterOrEqual},
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

	log.Fatal("[ Line ", Courant().NbLigne, ", Pos ", posToken%Courant().NbLigne, " ] Erreur : operateur ", typeWanted, " non trouvé.")
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

// Init : Initialise la liste
func Init(afterLexer []token.Token) {
	tokenTab = afterLexer
	posToken = 0
}

// Parser : Main function of this package
func Parser(afterLexer []token.Token) Noeud {
	tokenTab = afterLexer
	var mainNoeud Noeud = Fonction()

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

func Courant() token.Token {
	if posToken < len(tokenTab) {
		return tokenTab[posToken]
	}
	return token.Token{token.EOF, "", 0, -1}
}

func ajouterEnfant(parent Noeud, childs ...Noeud) Noeud {
	parent.Fils = append(parent.Fils, childs...)
	return parent
}

func verifier(typeCheck token.TokenType) bool {
	if Courant().DataType == typeCheck {
		avancer()
		return true
	}
	return false
}

func accepter(typeCheck token.TokenType) {
	if Courant().DataType != typeCheck {
		log.Fatal("[ Line ", Courant().NbLigne, ", Pos ", posToken%Courant().NbLigne, " ] Erreur accepter : type - ", typeCheck, " attendu, ", Courant().DataType, " (", Courant().ValeurString, ") reçu.")
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
		N = nouveauNoeud(NoeudSubUnaire, Courant().NbLigne)
		A := expression(getPrio(token.MinusUnaire).prioD)
		N = ajouterEnfant(N, A)
		return N
	} else if verifier(token.OperatorMult) {
		N = nouveauNoeud(NoeudIndirection, Courant().NbLigne)
		A := expression(getPrio(token.Pointeur).prioD)
		N = ajouterEnfant(N, A)
		return N
	} else if verifier(token.Not) {
		N = nouveauNoeud(NoeudNot, Courant().NbLigne)
		A := expression(getPrio(token.Not).prioD)
		N = ajouterEnfant(N, A)
		return N
	} else if Courant().DataType == token.Constant {
		N = nouveauNoeud(NoeudConst, Courant().NbLigne)
		N.ValeurEntiere = Courant().ValeurInt
		avancer()
		return N
	} else if Courant().DataType == token.Ident {
		T := Courant()
		avancer()
		if verifier(token.ParentheseOuvrante) {
			N = nouveauNoeud(NoeudAppel, T.NbLigne)
			N.ValeurString = T.ValeurString
			for Courant().DataType != token.ParentheseFermante {
				N = ajouterEnfant(N, expression(0))
				if !verifier(token.Virgule) {
					break
				}
			}
			accepter(token.ParentheseFermante)
			return N
		} else {
			N = nouveauNoeud(NoeudRef, T.NbLigne)
			N.ValeurString = T.ValeurString
			return N
		}
	}

	log.Fatal("[ Line ", Courant().NbLigne, ", Pos ", posToken, " ] Erreur : atome - ", token.Constant, " attendu, ", Courant().DataType, " reçu.")
	return N
}

func expression(prioMin int) Noeud {
	var N Noeud = atome()
	if estPrio(Courant().DataType) {
		for getPrio(Courant().DataType).prioG > prioMin {
			var op operation = getPrio(Courant().DataType)
			var line int = Courant().NbLigne
			avancer()
			var A Noeud = expression(op.prioD)
			var T Noeud = nouveauNoeud(op.typeOfNoeud, line)
			T = ajouterEnfant(T, N, A)
			N = T
			if !estPrio(Courant().DataType) {
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
		N = nouveauNoeud(NoeudDebug, Courant().NbLigne)
		N = ajouterEnfant(N, E)
		return N

	} else if verifier(token.KeywordSend) {
		E := expression(0)
		accepter(token.PointVirgule)
		N = nouveauNoeud(NoeudSend, Courant().NbLigne)
		N = ajouterEnfant(N, E)
		return N

	} else if verifier(token.LeftBrace) {
		N = nouveauNoeud(NoeudBlock, Courant().NbLigne)
		for !verifier(token.RightBrace) {
			N = ajouterEnfant(N, instruction())
		}
		return N

	} else if verifier(token.KeywordIf) {
		accepter(token.ParentheseOuvrante)
		E1 := expression(0)
		accepter(token.ParentheseFermante)
		I1 := instruction()
		N = nouveauNoeud(NoeudTest, Courant().NbLigne)
		N = ajouterEnfant(N, E1, I1)
		if verifier(token.KeywordElse) {
			I2 := instruction()
			N = ajouterEnfant(N, I2)
		}
		return N

	} else if verifier(token.KeywordWhile) {
		accepter(token.ParentheseOuvrante)
		E1 := expression(0)
		accepter(token.ParentheseFermante)
		I1 := instruction()
		N := nouveauNoeud(NoeudLoop, Courant().NbLigne)
		T := nouveauNoeud(NoeudTest, Courant().NbLigne)
		B := nouveauNoeud(NoeudBreak, Courant().NbLigne)
		T = ajouterEnfant(T, E1, I1, B)
		N = ajouterEnfant(N, T)
		return N

	} else if verifier(token.KeywordFor) {
		accepter(token.ParentheseOuvrante)
		E1 := expression(0)
		accepter(token.PointVirgule)
		E2 := expression(0)
		accepter(token.PointVirgule)
		E3 := expression(0)
		accepter(token.ParentheseFermante)
		I1 := instruction()

		N := nouveauNoeud(NoeudBlock, Courant().NbLigne)

		D1 := nouveauNoeud(NoeudDrop, Courant().NbLigne)
		L := nouveauNoeud(NoeudLoop, Courant().NbLigne)
		N = ajouterEnfant(N, D1, L)

		D1 = ajouterEnfant(D1, E1)
		C := nouveauNoeud(NoeudTest, Courant().NbLigne)
		L = ajouterEnfant(L, C)

		B2 := nouveauNoeud(NoeudBlock, Courant().NbLigne)
		BRAK := nouveauNoeud(NoeudBreak, Courant().NbLigne)
		C = ajouterEnfant(C, E2, B2, BRAK)

		D2 := nouveauNoeud(NoeudDrop, Courant().NbLigne)
		B2 = ajouterEnfant(B2, I1, D2)

		D2 = ajouterEnfant(D2, E3)

		return N

	} else if verifier(token.KeywordInt) {
		N = nouveauNoeud(NoeudDec, Courant().NbLigne)
		N.ValeurString = Courant().ValeurString
		avancer()
		accepter(token.PointVirgule)
		return N

	} else if verifier(token.KeywordBreak) {
		N = nouveauNoeud(NoeudBreak, Courant().NbLigne)
		//avancer()
		accepter(token.PointVirgule)
		return N

	} else if verifier(token.KeywordContinue) {
		N = nouveauNoeud(NoeudContinue, Courant().NbLigne)
		//avancer()
		accepter(token.PointVirgule)
		return N

	} else if verifier(token.Return) {
		E := expression(0)
		accepter(token.PointVirgule)
		N = nouveauNoeud(NoeudReturn, Courant().NbLigne)
		N = ajouterEnfant(N, E)
		return N

	} else {
		E1 := expression(0)
		accepter(token.PointVirgule)
		N = nouveauNoeud(NoeudDrop, Courant().NbLigne)
		N = ajouterEnfant(N, E1)
		return N

	}

}

//Fonction : Point d'entrée au parser
func Fonction() Noeud {
	accepter(token.KeywordInt)
	T := Courant()
	accepter(token.Ident)
	N := nouveauNoeud(NoeudFonction, Courant().NbLigne)
	N.ValeurString = T.ValeurString
	accepter(token.ParentheseOuvrante)

	for Courant().DataType != token.ParentheseFermante {
		if verifier(token.KeywordInt) {
			temp := Courant()
			accepter(token.Ident)
			D := nouveauNoeud(NoeudDec, temp.NbLigne)
			D.ValeurString = temp.ValeurString
			N = ajouterEnfant(N, D)
		} else {
			N = ajouterEnfant(N, expression(0))
		}
		if !verifier(token.Virgule) {
			break
		}
	}
	avancer()
	N = ajouterEnfant(N, instruction())
	return N
}
