package detectors

func isValidSLet(b byte) (isvl bool) {
	if b >= 97 && b <= 122 {
		isvl = true
	}
	return
}

func isValidCLet(b byte) (isvl bool) {
	if b >= 65 && b <= 90 {
		isvl = true
	}
	return
}

func isValidNum(b byte) (isvl bool) {
	if b >= 48 && b <= 57 {
		isvl = true
	}
	return
}

func Detect(s string) {
	if s != " " && s != "" {

	}
}
