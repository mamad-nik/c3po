package detectors

import (
	"fmt"
)

func Detect(s string) {
	if s != " " && s != "" {
		fmt.Printf("\t\t\t\tToken: %v, ", s)
		fmt.Printf("ASCII: %v\n", []byte(s))

	}
}
