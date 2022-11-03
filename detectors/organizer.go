package detectors

import (
	"fmt"

	"github.com/mamad-nik/lexer"
)

func Organize(token []string) {
	var state bool
	for j, x := range token {
		state = false
		for _, y := range lexer.Tab {
			for i := 0; i < len(y.Value); i++ {
				if x == y.Value[i] {
					fmt.Printf("num:%d, lexeme: %s, val: %s \n", j, x, y.Token)
					state = true
				}
			}

		}

		if state != true {
			Detect(x)
		}
	}
}
