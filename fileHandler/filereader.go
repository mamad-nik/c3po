package filehandler

import "os"

var token []string

func Read(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	temp := ""
	for i := 0; i < len(data)-1; i++ {
		WS := string(data[i]) == " " || string(data[i]) == "\n" || string(data[i]) == "\t"
		if WS {
			token = append(token, temp)
			temp = ""
		} else if string(data[i]) == ";" {
			token = append(token, temp, ";")
			temp = ""
		} else {
			temp = temp + string(data[i])
		}
	}
	return token
}
