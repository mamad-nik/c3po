package filehandler

import (
	"fmt"
	"os"

	"github.com/mamad-nik/lexer/detectors"
)

func Write(fileName string, ft []detectors.FinTable) {
	f, err := os.Create(fileName)
	defer f.Close()
	f.Chmod(0700)
	if err != nil {
		panic(err)
	}
	for _, v := range ft {
		s := fmt.Sprint(v.Index, ", ", v.Token, ", ", v.Value, ",\n")
		f.Write([]byte(s))
	}
}
