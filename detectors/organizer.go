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
		for _, y := range lexer.DefTable {
			for i := 0; i < len(y.Value); i++ {
				if x == y.Value[i] {
					ft = append(ft, FinTable{Index: j, Token: x, Value: y.Token})
					fmt.Printf("\nnum:%d, lexeme: %s, val: %s \n", j, x, y.Token)
					state = true
				}
			}
		}
		if state != true {
			fmt.Printf("\n-num:%d, lexeme: %s, \n", j, x)
			temp, err := Detect(x)
			if err != nil {
				fmt.Printf("compile failed because %s is %v\n", x, err)
				break
			} else {
				ft = append(ft, FinTable{Index: j, Token: x, Value: temp})
				fmt.Println(temp)
			}
		}
	}
	return ft
}
