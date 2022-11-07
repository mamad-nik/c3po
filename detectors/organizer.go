package detectors

import (
	"fmt"

	"github.com/mamad-nik/lexer"
)

type FinTable struct {
	Index int
	Token string
	Value string
}

func Organize(token []string, ft []FinTable) []FinTable {
	for j, x := range token {
		if elem, ok := lexer.Keyval[x]; ok {
			ft = append(ft, FinTable{Index: j, Token: x, Value: elem})
		} else {
			temp, err := Detect(x)
			if err != nil {
				fmt.Println(x)
				panic(err)
			} else {
				ft = append(ft, FinTable{Index: j, Token: x, Value: temp})
			}
		}
	}
	return ft
}
