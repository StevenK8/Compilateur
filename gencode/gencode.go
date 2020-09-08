package Gencode

import (
	parser "Compilateur/parser"
	"fmt"
)

var listOfAssembleurInstructions []string

func Gencode(Node parser.Noeud) {
	switch Node.TypeDeNoeud {

	case parser.NoeudConst:
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "push "+fmt.Sprint(Node.ValeurEntiere))
		break

	case parser.NoeudAdd:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "add")
		break

	case parser.NoeudSub:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "sub")
		break

	case parser.NoeudSubUnaire:
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "push 0")
		Gencode(Node.Fils[0])
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "sub")
		break

	case parser.NoeudMult:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "mul")
		break

	case parser.NoeudDiv:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "div")
		break

	case parser.NoeudMod:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "mod")
		break

	case parser.NoeudDebug:
		Gencode(Node.Fils[0])
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "dbg")
		break

	case parser.NoeudBlock:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		break

	case parser.NoeudDrop:
		Gencode(Node.Fils[0])
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "drop")
		break

	}
}

func Gen(Node parser.Noeud) []string {
	Gencode(Node)
	return listOfAssembleurInstructions
}
