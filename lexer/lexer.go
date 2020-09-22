package lexer

import (
	"regexp"
	"strconv"

	token "github.com/StevenK8/Compilateur/token"
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

		dataType, longueur := getOperator(data, charPos)

		if dataType != token.Ignore {
			i, _ := strconv.Atoi(string(data[charPos : charPos+longueur+1]))
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

	case "%":
		dataType = token.OperatorMod
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.EqualMod
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
		dataType = token.Not
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

	case ",":
		dataType = token.Virgule
		break

	case "<":
		dataType = token.LessThan
		break

	case ">":
		dataType = token.GreaterThan
		break
	default:
		return getIdent(data, charPos)
		break

	}
	return dataType, longueur
}

func getIdent(data []byte, charPos int) (token.TokenType, int) {
	var dataType token.TokenType
	dataType = token.Ident
	var longueur int

	if !checkMatchChar(`[a-zA-Z]`, string(data[charPos])) {
		return getNumber(data, charPos)
	}

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
			dataType = token.KeywordElse
		} else if charPos < len(data)-4 && string(data[charPos:charPos+5]) == "debug" && ((charPos < len(data)-5 && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+5]))) || (charPos < len(data)-4)) {
			dataType = token.Debug
		} else if charPos < len(data)-2 && string(data[charPos:charPos+3]) == "int" && ((charPos < len(data)-3 && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+3]))) || (charPos < len(data)-2)) {
			dataType = token.KeywordInt
		} else if charPos < len(data)-2 && string(data[charPos:charPos+3]) == "for" && ((charPos < len(data)-3 && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+3]))) || (charPos < len(data)-2)) {
			dataType = token.KeywordFor
		} else if charPos < len(data)-4 && string(data[charPos:charPos+5]) == "break" && ((charPos < len(data)-5 && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+5]))) || (charPos < len(data)-4)) {
			dataType = token.KeywordBreak
		} else if charPos < len(data)-7 && string(data[charPos:charPos+8]) == "continue" && ((charPos < len(data)-8 && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+8]))) || (charPos < len(data)-7)) {
			dataType = token.KeywordContinue
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

	if !checkMatchChar(`[0-9]`, string(data[charPos])) {
		return token.Ignore, 0
	}

	for longueur = 0; charPos < len(data)-1; charPos++ {
		if !checkMatchChar(`[0-9]`, string(data[charPos+1])) {
			break
		}
		longueur++
	}

	return dataType, longueur
}
