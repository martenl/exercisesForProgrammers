package main

import (
	"fmt"
	"os"
)

func countingCharacters() {
	input := ""
	for input == "" {
		fmt.Print("What is the input string? ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			os.Exit(-1)
		}
		if input == "" {
			fmt.Println("Please enter a string")
		}
	}
	noOfCharacters := len(input)
	fmt.Printf("%s has %d characters.", input, noOfCharacters)
}
