package main

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func main() {
	fmt.Print("Base 64 encoder\n")
	fmt.Print("Enter string: ")
	var input string
	fmt.Scanln(&input)
	encoded_str := base64Encode(input)
	fmt.Println(encoded_str)
}
