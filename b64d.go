package main

import (
	"encoding/base64"
	"fmt"
)

func base64Decode(str string) (string, bool) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", true
	}
	return string(data), false
}

func main() {
	fmt.Print("Base 64 decoder\n")
	fmt.Print("Enter string: ")
	var input string
	fmt.Scanln(&input)

	decoded_str, decode_err := base64Decode(input)
	if decode_err {
		fmt.Println("Decoding failed.")
	}
	fmt.Println(decoded_str)
}
