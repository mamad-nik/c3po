package lexer

type defEntity struct {
	Token string
	Value []string
}

var DefTable []defEntity = []defEntity{
	{
		"class",
		class,
	},
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
		"return",
		returnstmt,
	},
	{
		"curBrOp",
		curBrOp,
	},
	{
		"curBrCl",
		curBrCl,
	},
	{
		"BrOp",
		BrOp,
	},
	{
		"BrCl",
		BrCl,
	},
	{
		"protect",
		levelOfProtection,
	},
}
var class []string = []string{
	"class",
}
var levelOfProtection []string = []string{
	"public",
	"private",
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
var returnstmt []string = []string{
	"return",
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
var BrOp []string = []string{
	"[",
}
var BrCl []string = []string{
	"]",
}
