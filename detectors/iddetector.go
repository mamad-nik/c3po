package detectors

import "errors"

func isValidSLet(b byte) bool {
	return b >= 97 && b <= 122
}

func isValidCLet(b byte) bool {
	return b >= 65 && b <= 90
}

func isValidDig(b byte) bool {
	return b >= 48 && b <= 57
}
func isValidNum(b []byte) (iv bool) {
	for i := 0; i < len(b); i++ {
		iv = isValidDig(b[i])
	}
	return
}
func isValidLet(b byte) bool {
	return isValidCLet(b) || isValidSLet(b)

}

func isValidID(b []byte) (isvl bool) {
	var v1, v2 bool
	if isValidLet(b[0]) {
		v1 = true
	}
	if len(b) > 1 {
		for i := 1; i < len(b); i++ {
			if isValidLet(b[i]) || isValidDig(b[i]) {
				v2 = true
			}
		}
	} else {
		v2 = true
	}
	return v1 && v2
}

func Detect(s string) (string, error) {
	if s != " " && s != "" {
		if isValidID([]byte(s)) {
			return "ID", nil
		} else if isValidNum([]byte(s)) {
			return "NUM", nil
		} else {
			return "", errors.New("invalid type")
		}
	}
	return "", errors.New("compiler error")
}
