package main

import (
	"fmt"
	"os"

	"github.com/mamad-nik/lexer/detectors"
	filehandler "github.com/mamad-nik/lexer/fileHandler"
)

var token []string

var ft []detectors.FinTable

func main() {
	var input stringS
	fmt.Scan(&input)
	token = filehandler.Read(input)
	ft = detectors.Organize(token, ft)
	f, err := os.CreateTemp("test", "final.csv")
	if err != nil {
		panic(err)
	}
	filehandler.Write(f.Name(), ft)
}
