package main

import (
	"fmt"

	"github.com/mamad-nik/lexer/detectors"
	filehandler "github.com/mamad-nik/lexer/fileHandler"
)

var token []string

func main() {
	token = filehandler.Read("/home/mamad/testArea/testComp")
	fmt.Println(token)
	detectors.Organize(token)

}
