package reverse_string

func ReverseString(input string) (output string) {
	for i := len(input) - 1; i >= 0; i-- {
		res := (string(input[i]))
		output += res
	}
	return output
}

/*package main

import (
	"fmt"
)

func main() {
	s := "Hello world!"
	grim := ReverseString(s)
	fmt.Println(grim)
}
func ReverseString(input string) (output string) {
	for i := len(input) - 1; i >= 0; i-- {
		res := (string(input[i]))
		output += res
	}
	return output
}*/
