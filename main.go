package main

import (
	lexer "Compilateur/lexer"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	tokenTab = lexer.Lexer(data)
	for _, a := range tokenTab {
		println(strconv.Itoa(a.nbLigne) + " \t[" + string(a.valeurString) + "]\t" + string(a.dataType))
	}

	parser(tokenTab)

}
