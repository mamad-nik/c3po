package lexer

type Entity struct {
	Token string
	Value []string
}

var Tab []Entity = []Entity{
	{
		"keyword",
		keywords,
	},
	{
		"prim",
		primitive,
	},
	{
		"relop",
		relop,
	},
	{
		"logop",
		logop,
	},
	{
		"ariop",
		ariop,
	},
	{
		"eol",
		eol,
	},
	{
		"assign",
		assign,
	},
	{
		"openpar",
		openpar,
	},
	{
		"closepar",
		closepar,
	},
	{
		"curBrOp",
		curBrOp,
	},
	{
		"curBrCl",
		curBrCl,
	},
}
var kes [][]byte = [][]byte{
	[]byte("mamad"),
}

var keywords []string = []string{
	"if",
	"while",
}
var primitive []string = []string{
	"void",
	"int",
	"boolean",
}
var relop []string = []string{
	">",
	"<",
	"==",
	"<=",
	">=",
	"!=",
}
var logop []string = []string{
	"&&",
	"||",
	"!",
}

var ariop []string = []string{
	"+",
	"-",
	"*",
	"/",
}
var assign []string = []string{
	"=",
}
var eol []string = []string{
	";",
}
var openpar []string = []string{
	"(",
}

var closepar []string = []string{
	")",
}
var curBrOp []string = []string{
	"{",
}
var curBrCl []string = []string{
	"}",
}
