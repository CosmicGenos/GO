package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("Please enter your name:")
	reader := bufio.NewReader(os.Stdin)
	//go dont have built-in function to read input from the user, so we use bufio.NewReader to read input from standard
	// input (os.Stdin). The ReadString method reads until it encounters a newline character ('\n').
	// and go dont have  error handling, so we ignore the error returned by ReadString method using the blank identifier (_).
	output, _ := reader.ReadString('\n')
	fmt.Println("Hello, " + output)

	fmt.Println("Please enter your age:")
	ageReader := bufio.NewReader(os.Stdin)
	ageInput, _ := ageReader.ReadString('\n')
	fmt.Println("You are " + ageInput + " years old.")

	addoneage,err :=  strconv.ParseFloat(strings.TrimSpace(ageInput), 32)
	if err != nil {
		fmt.Println("Error parsing age:", err)
		return
	}
	fmt.Println("You are " + fmt.Sprint(addoneage) + " years old.")

}
