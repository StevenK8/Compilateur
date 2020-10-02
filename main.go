package main

/*
	https://perso.limsi.fr/lavergne/
*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/StevenK8/Compilateur/gencode"
	"github.com/StevenK8/Compilateur/lexer"
	"github.com/StevenK8/Compilateur/parser"
	"github.com/StevenK8/Compilateur/semantique"
	"github.com/StevenK8/Compilateur/token"
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

	var g []string

	g = append(g, ".start", "resn "+fmt.Sprint(semantique.NbSlot))

	for parser.Courant().DataType != token.EOF {
		N := parser.Fonction()
		parser.PrintNoeud(N, 0)
		N = semantique.Sem(N)
		g = append(g, gencode.Gen(N)...)
	}

	g = append(g, "halt")

	f, err := os.Create("test.out")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, gen := range g {
		_, err = f.WriteString(gen + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}
