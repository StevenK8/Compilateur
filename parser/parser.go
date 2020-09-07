package parser

import (
	token "Compilateur/token"
	"log"
)

type typeNoeud string

const (
	noeudPlus				typeNoeud = "Noeud_Add"
	noeudMinus				typeNoeud = "Noeud_Sub"
	noeudOperatorMult       typeNoeud = "Noeud_Mult"
	noeudOperatorDiv		typeNoeud = "Noeud_Div"
	noeudOperatorModul	   	typeNoeud = "Noeud_Mod"
	noeudParentheseOuvrante typeNoeud = "Noeud_Open_Paren"
	noeudParentheseFermante typeNoeud = "Noeud_Close_Paren"
	noeudLeftBrace          typeNoeud = "Noeud_Open_Brace"
	noeudRightBrace         typeNoeud = "Noeud_Close_Brace"
	noeudPointVirgule       typeNoeud = "Noeud_Semicolon"
	noeudconstant           typeNoeud = "Noeud_Number"
)

type priorite struct {
    prioBas		int
    prioHaut 	int
    typeOfNoeud	typeNoeud
}

const prioBase []priorite { 
		priorite {10, 11, operatorPlus},
		priorite {10, 11, operatorMinus},
		priorite {20, 21, operatorMult},
		priorite {20, 21, operatorDiv},
		priorite {20, 21, operatorModul},
		priorite {60, 61, noeudParentheseOuvrante},
}

type noeud struct {
	filsG         *noeud
	filsD         *noeud
	valeurString  string
	valeurEntiere int
	nbLigne       int
	typeDeNoeud   string
}

var tokenTab []token.Token
var posToken int

func parser([]token.Token) {

}

func avancer() {
	posToken++
}

func courant() token.Token {
	return tokenTab[posToken]
}

func ajouterEnfant(t noeud, n noeud, a noeud) noeud {
	t.filsG = &n
	t.filsD = &a
	return t
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
		log.Fatal("accepter")
	}
	avancer()
}

func a() noeud {
	var N noeud
	if verifier(token.ParentheseOuvrante) {
		N = e(0)
		accepter(token.ParentheseFermante)
	} else if verifier(token.ParentheseFermante) {
		N = e(0)
		accepter(token.ParentheseOuvrante)
	}
	return N
}

func e(valeur int) noeud {
	//TODO
	return noeud{nil, nil, "", valeur, 0, ""}
}

// func expression(prioMin int) noeud {
// 	N := a()
// 	for OP[courant().dataType].prio > prioMin {
// 		op = OP[courant().dataType]
// 		avancer()
// 		A := e(op.prio)
// 		T := noeud{nil, nil, "", 0, 0, op.typeDeNoeud}
// 		T = ajouterEnfant(T, N, A)
// 		N = T
// 	}
// 	return N
// }

func atome() noeud {
	if courant().DataType == token.Constant {
		N := noeud{nil, nil, courant().ValeurString, courant().ValeurInt, courant().NbLigne, "type"}
		avancer()
		return N
	}
	log.Fatal("atome")
	return noeud{nil, nil, "", 0, 0, ""}
}
