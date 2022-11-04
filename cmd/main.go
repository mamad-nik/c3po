package main

import (
	"fmt"

	"github.com/mamad-nik/lexer/detectors"
	filehandler "github.com/mamad-nik/lexer/fileHandler"
)

var token []string

<<<<<<< HEAD

=======
>>>>>>> parent of 18e7d54 (started implementing saving architecture)
func main() {
	var ft &[]lexer.FinTable

	token = filehandler.Read("/home/mamad/testArea/testComp")
	fmt.Println(token)
<<<<<<< HEAD
	detectors.Organize(token, ft)
	fmt.Println(*ft)
=======
	detectors.Organize(token)

>>>>>>> parent of 18e7d54 (started implementing saving architecture)
}
