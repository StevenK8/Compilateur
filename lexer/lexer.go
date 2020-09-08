package lexer

import (
	"regexp"
	"strconv"

	token "Compilateur/token"
)

func Lexer(data []byte) []token.Token {

	var tokenTab []token.Token
	numOfLine := 1

	for charPos := 0; charPos < len(data); charPos++ {

		currentChar := string(data[charPos])

		changeLine, err := regexp.MatchString(`\n`, currentChar)
		if err != nil {
			println(err)
		}
		if changeLine {
			numOfLine++
		}

		if checkMatchChar("[<{(+*;=!/&|)}>]", currentChar) {
			dataType, longueur := getOperator(data, charPos)
			tokenTab = append(tokenTab, token.Token{dataType, string(data[charPos : charPos+longueur+1]), 0, numOfLine})
			charPos += longueur
		} else if checkMatchChar("[a-zA-Z]", currentChar) {
			dataType, longueur := getIdent(data, charPos)
			tokenTab = append(tokenTab, token.Token{dataType, string(data[charPos : charPos+longueur+1]), 0, numOfLine})
			charPos += longueur
		} else if checkMatchChar("[0-9]", currentChar) {
			dataType, longueur := getNumber(data, charPos)
			i, err := strconv.Atoi(string(data[charPos : charPos+longueur+1]))
			if err != nil {
				println(err)
			}
			tokenTab = append(tokenTab, token.Token{dataType, string(data[charPos : charPos+longueur+1]), i, numOfLine})
			charPos += longueur
		}

	}
	return tokenTab
}

func checkMatchChar(regex string, char string) bool {
	matched, err := regexp.MatchString(regex, char)
	if err != nil {
		println(err)
	}

	return matched
}

func getOperator(data []byte, charPos int) (token.TokenType, int) {
	var dataType token.TokenType
	longueur := 0
	switch string(data[charPos]) {
	case "+":
		dataType = token.OperatorPlus
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.EqualAdd
			} else if string(data[charPos+1]) == "+" {
				longueur++
				dataType = token.Increment
			}
		}

		break
	case "-":
		dataType = token.OperatorMinus
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.EqualSub
			}
		}

		break
	case "*":
		dataType = token.OperatorMult
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.EqualMult
			} else if string(data[charPos+1]) == "*" {
				longueur++
				dataType = token.Pow
			}
		}
		break

	case "/":
		dataType = token.OperatorDiv
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.EqualDiv
			}
		}
		break

	case "=":
		dataType = token.Equal
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.Equalequal
			}
		}
		break

	case "!":
		dataType = token.Equal
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.NotEqual
			}
		}
		break

	case "&":
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "&" {
				longueur++
				dataType = token.And
			}
		}
		break

	case "|":
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "|" {
				longueur++
				dataType = token.Or
			}
		}
		break

	case "(":
		dataType = token.ParentheseOuvrante
		break
	case ")":
		dataType = token.ParentheseFermante
		break
	case "{":
		dataType = token.LeftBrace
		break
	case "}":
		dataType = token.RightBrace
		break
	case ";":
		dataType = token.PointVirgule
		break
	case "<":
		dataType = token.LessThan
		break
	case ">":
		dataType = token.GreaterThan
		break
	}
	return dataType, longueur
}

func getIdent(data []byte, charPos int) (token.TokenType, int) {
	var dataType token.TokenType
	dataType = token.Word
	var longueur int

	for longueur = 0; charPos < len(data)-1; charPos++ {
		if charPos < len(data)-1 && string(data[charPos:charPos+2]) == "if" && ((charPos < len(data)-2 && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+2]))) || (charPos < len(data)-1)) {
			dataType = token.KeywordIf
		} else if charPos < len(data)-4 && string(data[charPos:charPos+5]) == "while" && ((charPos < len(data)-5 && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+5]))) || (charPos < len(data)-4)) {
			dataType = token.KeywordWhile
		} else if charPos < len(data)-3 && string(data[charPos:charPos+4]) == "true" && ((charPos < len(data)-4 && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+4]))) || (charPos < len(data)-3)) {
			dataType = token.BooleanTrue
		} else if charPos < len(data)-4 && string(data[charPos:charPos+5]) == "false" && ((charPos < len(data)-5 && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+5]))) || (charPos < len(data)-4)) {
			dataType = token.BooleanFalse
		} else if charPos < len(data)-3 && string(data[charPos:charPos+4]) == "else" && ((charPos < len(data)-4 && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+4]))) || (charPos < len(data)-3)) {
			dataType = token.BooleanTrue
		}

		if !checkMatchChar(`[a-zA-Z]`, string(data[charPos+1])) {
			break
		}

		longueur++
	}

	return dataType, longueur
}

func getNumber(data []byte, charPos int) (token.TokenType, int) {
	var dataType token.TokenType
	dataType = token.Constant
	var longueur int

	for longueur = 0; charPos < len(data)-1; charPos++ {
		if !checkMatchChar(`[0-9]`, string(data[charPos+1])) {
			break
		}
		longueur++
	}

	return dataType, longueur
}
