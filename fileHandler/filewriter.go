package filehandler

import (
	"fmt"
	"os"

	"github.com/mamad-nik/lexer/detectors"
)

func Write(ft []detectors.FinTable, f *os.File) {
	for _, v := range ft {
		s := fmt.Sprint(v.Index, ", ", v.Token, ", ", v.Value, ",\n")
		f.Write([]byte(s))
	}
}
