package main

import "fmt"

func checkPalindrome(str string) bool {

	paly := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		paly = append(paly, str[i])
	}
	return str == string(paly)

}

func main() {

	fmt.Printf("%+v\n", checkPalindrome("ana"))
	fmt.Printf("%+v\n", checkPalindrome("test"))

}
