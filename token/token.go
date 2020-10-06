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
	Not                = "!"
	LessThan           = "<"
	LessOrEqual        = "<="
	GreaterThan        = ">"
	GreaterOrEqual     = ">="
	Equalequal         = "=="
	Equal              = "Equal"       // =
	OperatorPlus       = "Add"         // +
	OperatorMinus      = "Sub"         // -
	OperatorMult       = "Mult"        // *
	OperatorDiv        = "Div"         // /
	OperatorMod        = "Mod"         // %
	ParentheseOuvrante = "Open_Paren"  // (
	ParentheseFermante = "Close_Paren" // )
	LeftBrace          = "Open_Brace"  // {
	RightBrace         = "Close_Brace" // }
	PointVirgule       = "Semicolon"   // ;
	Virgule            = "Virgule"     // ,
	Constant           = "Number"
	Ident              = "Ident"
	KeywordIf          = "If"
	KeywordElse        = "Else"
	KeywordWhile       = "While"
	KeywordFor         = "For"
	KeywordBreak       = "Break"
	KeywordContinue    = "Continue"
	KeywordInt         = "int"
	KeywordSend        = "Send"
	BooleanTrue        = "True"
	BooleanFalse       = "False"
	And                = "And"
	Or                 = "Or"
	EOF                = "EOF"
	MinusUnaire        = "-u"
	Pointeur           = "*u"
	Debug              = "Debug"
	Ignore             = "Ignored"
	Return             = "Return"
)

type Token struct {
	DataType     TokenType
	ValeurString string
	ValeurInt    int
	NbLigne      int
}
