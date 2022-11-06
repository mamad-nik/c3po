package main

import (
	"github.com/mamad-nik/lexer/detectors"
	filehandler "github.com/mamad-nik/lexer/fileHandler"
)

var token []string

var ft []detectors.FinTable

func main() {

	token = filehandler.Read("/home/mamad/go/src/lexer/test/testComp.txt")
	ft = detectors.Organize(token, ft)
	filehandler.Write("final.csv", ft)
}
