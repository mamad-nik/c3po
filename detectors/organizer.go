package detectors

import (
	"fmt"

	"github.com/mamad-nik/lexer"
)

func Organize(token []string) {
	for j, x := range token {
		state := false
		for _, y := range lexer.DefTable {
			for i := 0; i < len(y.Value); i++ {
				if x == y.Value[i] {
					fmt.Printf("num:%d, lexeme: %s, val: %s \n", j, x, y.Token)
					state = true
				}
			}
		}
		if state != true {
			fmt.Printf("num:%d, lexeme: %s, \n", j, x)
			Detect(x)
		}
	}
}
