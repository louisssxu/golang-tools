package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func readFile(path string) string {
	data, err := os.ReadFile(path)
	check(err)
	return string(data)
}

func countChar (text string) int{
	return len(text)
}
func countLine (text string) int{
	count := 0
	for _, r := range text {
		if r == '\n'{ count++}
	}
	return count
}
func countWord (text string) int{
	count := 0
	isPrecededByChar := false
	for _, r := range text {
		if r == '\n' || r == '\t' || r == ' ' ||  r == '\r'{
			isPrecededByChar = false
		}else{
			if !isPrecededByChar {
				count++
			}
			isPrecededByChar = true
		}
	}
	return count
}


func main()  {
	args := os.Args[1:]
	argsLen := len(args)
 
	charFlag := false
	wordFlag := false
	lineFlag := false
	allFlag := true

	var targets []string

	
	for i:=0 ; i < argsLen ; i++ {

		arg := args[i]
		if targets != nil {
			targets = append(targets, arg)
			continue
		}
		firstRune, _ := utf8.DecodeRuneInString(arg)
		if (firstRune == '-'){
			for i , r := range arg{
				if i == 0 {continue}
				switch r {
				case 'l':
					lineFlag = true
					allFlag = false
				case 'w':
					wordFlag = true
					allFlag = false
				case 'c':
					charFlag = true
					allFlag = false
				default:
					fmt.Println("ccwc: illegal option -- ", string(r))
					panic(r)
				}
			}
			continue
		}
		targets = append(targets, arg)
	}

	


	for _, tar := range targets {
		text := readFile(tar)
		if allFlag {
			fmt.Print(countLine(text), "\t")
			fmt.Print(countWord(text), "\t")
			fmt.Print(countChar(text), "\t")
		}else{
			if lineFlag {
				fmt.Print(countLine(text), "\t")
			}
			if wordFlag{
				fmt.Print(countWord(text), "\t")
			}
			if charFlag{
				fmt.Print(countChar(text), "\t")
			}
		}
		fmt.Print(tar)
		fmt.Println()
	}

}