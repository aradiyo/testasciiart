package main

import (
	//"bufio"
	"fmt"
	"os"
	//"strings"
)

func main() {
  if len(os.Args) == 1 {
		fmt.Println("Please enter your text to be converted into ascii-art!")
    return
	}
  
	input := os.Args[1]
//	var font string
// font := os.Args[2]

//  if len(os.Args) >= 3 {
//    font = os.Args[2]
//  } else {
//    font = "standard"
//  }

  if !isNewLine(input) {
    fmt.Println("there is no new line")
  } else {
    fmt.Println("New Line detected!")
  }

  
	

	if !isValidAscii(input) {
		fmt.Println("Invalid input! Please enter valid ascii input ")
	 }

	fmt.Println(input)
}

func isValidAscii(input string) bool {
	for _, x := range input {
		if x < 32 || x > 126 {
			return false
		}
	}
	return true
}
func isNewLine(input string) bool {
	for i, x := range input {
		if i > 0 && x == 'n' && input[i-1] == '\\' {
			return true
		}
	}

	return false
}
