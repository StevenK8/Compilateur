package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	gencode "Compilateur/gencode"
	lexer "Compilateur/lexer"
	parser "Compilateur/parser"
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
	g := gencode.Gen(p)
	for _, instruction := range g {
		fmt.Println(instruction)
	}

}
