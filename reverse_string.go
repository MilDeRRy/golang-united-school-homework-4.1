package reverse_string

func ReverseString(input string) (output string) {
	em := []rune(input)
	for i := len(em) - 1; i >= 0; i-- {
		res := (string(em[i]))
		output += res
	}
	return output
}

/*package main

import (
	"fmt"
)

func main() {
	s := "Привет"
	grim := ReverseString(s)
	fmt.Println(grim)
}
func ReverseString(input string) (output string) {
	em := []rune(input)
	for i := len(em) - 1; i >= 0; i-- {
		res := (string(em[i]))
		output += res
	}
	//ty := []byte(input)
	//fmt.Println(ty)
	//bytes := []byte{209, 130, 208, 181, 208, 178, 208, 184, 209, 128, 208, 159}
	//str := string(bytes)
	//fmt.Println(str)
	return output
}*/
