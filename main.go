package main

/*
	https://perso.limsi.fr/lavergne/
*/

import (
	"fmt"
	"io/ioutil"
	"strconv"

	//"Compilateur/gencode"
	"Compilateur/gencode"
	"Compilateur/lexer"
	"Compilateur/parser"
	"Compilateur/semantique"
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

	p := parser.Parser(tokenTab)
	p = semantique.Sem(p)
	parser.PrintNoeud(p, 0)

	gencode.ListOfAssembleurInstructions = append(gencode.ListOfAssembleurInstructions, ".start")
	gencode.ListOfAssembleurInstructions = append(gencode.ListOfAssembleurInstructions, "resn "+fmt.Sprint(semantique.NbSlot))

	g := gencode.Gen(p)
	for _, instruction := range g {
		fmt.Println(instruction)
	}

}
