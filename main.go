package main

/*
	https://perso.limsi.fr/lavergne/
*/

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/StevenK8/Compilateur/gencode"
	"github.com/StevenK8/Compilateur/lexer"
	"github.com/StevenK8/Compilateur/parser"
	"github.com/StevenK8/Compilateur/semantique"
	"github.com/StevenK8/Compilateur/token"
)

func main() {
	fileName := flag.String("file", "", "path of input file")
	fileNameOut := flag.String("o", "", "path of output file")
	boolPtr := flag.Bool("h", false, "a bool")

	flag.Parse()

	if *boolPtr || *fileName == "" {
		fmt.Println("Utilisation du Compilateur:\n" +
			"./Compilateur -file='test.txt' -o='test.out'")
		os.Exit(0)
	}

	data, err := ioutil.ReadFile(*fileName)
	if err != nil {
		fmt.Println("File reading error:", err)
		return
	}

	tokenTab := lexer.Lexer(data)
	for _, a := range tokenTab {
		fmt.Println(strconv.Itoa(a.NbLigne) + " \t[" + string(a.ValeurString) + "]\t" + string(a.DataType))
	}

	semantique.DebutBlock()
	parser.Init(tokenTab)

	var g []string

	for parser.Courant().DataType != token.EOF {
		N := parser.Fonction()
		parser.PrintNoeud(N, 0)
		N = semantique.Sem(N)
		g = append(g, gencode.Gen(N)...)
	}
	// g = append(g, ".start", "resn "+fmt.Sprint(semantique.NbSlot))

	g = append(g, ".start", "prep main", "call 0", "halt")

	var outPath string
	if *fileNameOut == "" {
		outPath = strings.Split(*fileName, ".")[0] + ".out"
	} else {
		outPath = *fileNameOut
	}

	f, err := os.Create(outPath)

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
