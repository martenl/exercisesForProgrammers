package main

import (
	"fmt"
	"os"
)

func sayingHello() {
	fmt.Print("What is your name? ")
	name := ""
	_, err := fmt.Scanln(&name)
	if err != nil {
		os.Exit(-1)
	}
	fmt.Printf("Hello, %s, nice to meet you", name)
}
