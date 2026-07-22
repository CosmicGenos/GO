package main

import "fmt"
import "bufio"
import "os"

func main() {
	fmt.Println("Hello, World!")
	
	fmt.Println("Please enter your name:")
	reader := bufio.NewReader(os.Stdin)
	//go dont have built-in function to read input from the user, so we use bufio.NewReader to read input from standard 
	// input (os.Stdin). The ReadString method reads until it encounters a newline character ('\n').
	// and go dont have  error handling, so we ignore the error returned by ReadString method using the blank identifier (_).
	output, _ := reader.ReadString('\n') 
	fmt.Println("Hello, " + output)
	
	
}