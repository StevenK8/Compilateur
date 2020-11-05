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

func compileRuntime() {
	runtime, err := ioutil.ReadFile("runtime.h")
	if err != nil {
		fmt.Println("Runtime reading error:", err)
		return
	}

	compile(runtime) // Compilation du runtime
}

func main() {
	fileName := flag.String("file", "test.txt", "path of input file")

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

	compileRuntime()

	compile(data) // Compilation du code source

	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})

	writeOutput(*fileName)
}

func writeOutput(fileName string) {
	fileNameOut := flag.String("o", "", "path of output file")

	var outPath string
	if *fileNameOut == "" {
		outPath = strings.Split(fileName, ".")[0] + ".out"
	} else {
		outPath = *fileNameOut
	}

	f, err := os.Create(outPath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	g := gencode.GetGenList()

	for _, gen := range g {
		_, err = f.WriteString(gen + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func compile(data []byte) {
	tokenTab := lexer.Lexer(data)
	for _, a := range tokenTab {
		fmt.Println(strconv.Itoa(a.NbLigne) + " \t[" + string(a.ValeurString) + "]\t" + string(a.DataType))
	}

	semantique.DebutBlock()
	parser.Init(tokenTab)

	for parser.Courant().DataType != token.EOF {
		N := parser.Fonction()
		parser.PrintNoeud(N, 0)
		N = semantique.Sem(N)
		gencode.Gen(N)
	}
}
