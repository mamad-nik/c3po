package lexer

type defEntity struct {
	Token string
	Value []string
}

var Keyval = map[string]string{
	"class":   "class",
	"if":      "keyword",
	"while":   "keyword",
	"public":  "protect",
	"private": "protect",
	"void":    "prim",
	"int":     "prim",
	"boolean": "prim",
	"static":  "static",
	"this":    "ref",
	"=":       "assign",
	">":       "relop",
	"<":       "relop",
	"==":      "relop",
	"<=":      "relop",
	">=":      "relop",
	"!=":      "relop",
	"&&":      "logop",
	"||":      "logop",
	"!":       "logop",
	"+":       "ariop",
	"-":       "ariop",
	"*":       "ariop",
	"/":       "ariop",
	"return":  "returnstmt",
	";":       "EOL",
	"(":       "openpar",
	")":       "closepar",
	"{":       "opcurbr",
	"}":       "clcurbr",
	"[":       "opbr",
	"]":       "clbr",
	"\"":      "qouma",
}
