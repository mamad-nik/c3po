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
		state := false
		for key, val := range lexer.Keyval {
			if key == x {
				ft = append(ft, FinTable{Index: j, Token: x, Value: val})
			}
		}

		if !state {
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
