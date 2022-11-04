package main

import (
	"fmt"

	"github.com/mamad-nik/lexer"
	"github.com/mamad-nik/lexer/detectors"
	filehandler "github.com/mamad-nik/lexer/fileHandler"
)

var token []string

var ft []lexer.FinTable

func main() {
	token = filehandler.Read("/home/mamad/testArea/testComp")
	fmt.Println(token)
	ft = detectors.Organize(token)
	fmt.Println(ft)
}
