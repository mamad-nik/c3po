package main

import (
	"github.com/mamad-nik/lexer/detectors"
	filehandler "github.com/mamad-nik/lexer/fileHandler"
)

var token []string

var ft []detectors.FinTable

func main() {
	var input string
	fmt.scan(&input)
	token = filehandler.Read(input)
	ft = detectors.Organize(token, ft)
	filehandler.Write("final.txt", ft)
}
