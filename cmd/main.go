package main

import (
	"os"

	"github.com/mamad-nik/lexer/detectors"
	filehandler "github.com/mamad-nik/lexer/fileHandler"
)

var token []string

var ft []detectors.FinTable

func main() {
	token = filehandler.Read("/home/mamad/testArea/testComp")
	ft = detectors.Organize(token, ft)
	f, err := os.Create("/home/mamad/testArea/final")
	defer f.Close()
	f.Chmod(0700)
	if err != nil {
		panic(err)
	}
	filehandler.Write(ft, f)
}
