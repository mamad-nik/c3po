package filehandler

import (
	"fmt"
	"os"
)

var token []string

func appender(arr *[]string, tok string) {
	if tok != "" {
		*arr = append(*arr, tok)
	}
}

func Read(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	temp := ""
	for i := 0; i < len(data)-1; i++ {
		WS := string(data[i]) == " " || string(data[i]) == "\n" || string(data[i]) == "\t"
		Para := string(data[i]) == "(" || string(data[i]) == ")" || string(data[i]) == "[" ||
			string(data[i]) == "]" || string(data[i]) == "{" || string(data[i]) == "}" ||
			string(data[i]) == "\"" || string(data[i]) == "," || string(data[i]) == "+" ||
			string(data[i]) == "-" || string(data[i]) == "*" || string(data[i]) == "/"

		if WS {
			appender(&token, temp)
			temp = ""
		} else if string(data[i]) == ";" {
			appender(&token, temp)
			appender(&token, ";")
			temp = ""
		} else if Para {
			appender(&token, temp)
			appender(&token, string(data[i]))
			temp = ""
		} else {
			temp = temp + string(data[i])
		}
	}
	fmt.Println(token)
	return token
}
