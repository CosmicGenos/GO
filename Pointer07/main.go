package main

import "fmt"

func main() {
	fmt.Println("hello ")

	var ptr *int

	var a int = 10

	ptr = &a

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of ptr:", ptr)
	fmt.Println("Value pointed to by ptr:", *ptr)
}