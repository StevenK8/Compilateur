package main

/*
	https://perso.limsi.fr/lavergne/
*/

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"Compilateur/gencode"
	"Compilateur/lexer"
	"Compilateur/parser"
	"Compilateur/semantique"
	"Compilateur/token"
)

func main() {

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	tokenTab := lexer.Lexer(data)
	for _, a := range tokenTab {
		fmt.Println(strconv.Itoa(a.NbLigne) + " \t[" + string(a.ValeurString) + "]\t" + string(a.DataType))
	}

	semantique.DebutBlock()
	parser.Init(tokenTab)

	gencode.ListOfAssembleurInstructions = append(gencode.ListOfAssembleurInstructions, ".start")
	gencode.ListOfAssembleurInstructions = append(gencode.ListOfAssembleurInstructions, "resn "+fmt.Sprint(semantique.NbSlot))

	var g []string

	for parser.Courant().DataType != token.EOF {
		N := parser.Fonction()
		parser.PrintNoeud(N, 0)
		N = semantique.Sem(N)
		for _, tab := range gencode.Gen(N) {
			g = append(g, tab)
		}
	}

	fmt.Println(g)
}
