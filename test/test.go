package main

import (
	"fmt"
	"unicode/utf8"
)



func main(){
	var a rune

	var b *rune

	fmt.Println(b)

	b = &a

	fmt.Println(a, b, *b)

	fmt.Println(len("abcd"))
	fmt.Println(utf8.RuneCountInString("abcd"))
}