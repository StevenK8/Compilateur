package token

type TokenType string

const (
	EqualAdd           = "+="
	EqualSub           = "-="
	EqualMult          = "*="
	EqualDiv           = "/="
	EqualMod           = "%="
	Increment          = "++"
	Pow                = "**"
	NotEqual           = "!="
	LessThan           = "<"
	GreaterThan        = ">"
	Equalequal         = "=="
	Equal              = "Equal"
	OperatorPlus       = "Add"
	OperatorMinus      = "Sub"
	OperatorMult       = "Mult"
	OperatorDiv        = "Div"
	OperatorMod        = "Mod"
	ParentheseOuvrante = "Open_Paren"
	ParentheseFermante = "Close_Paren"
	LeftBrace          = "Open_Brace"
	RightBrace         = "Close_Brace"
	PointVirgule       = "Semicolon"
	Constant           = "Number"
	Word               = "Word"
	KeywordIf          = "If"
	KeywordElse        = "Else"
	KeywordWhile       = "While"
	BooleanTrue        = "True"
	BooleanFalse       = "False"
	And                = "And"
	Or                 = "Or"
	EOF                = "EOF"
)

type Token struct {
	DataType     TokenType
	ValeurString string
	ValeurInt    int
	NbLigne      int
}
