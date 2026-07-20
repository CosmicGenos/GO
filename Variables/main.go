package main

import "fmt"

func main(){
	var username string = "Kavindu"
	fmt.Println(username)

	var age int = 26

	fmt.Printf("My name is %s and I am %d years old.", username, age)

	var isloggedin bool = true

	if isloggedin {
		fmt.Println("User is logged in")
	}
	fmt.Println("User is not logged in")
	
	var u1 uint8 = 255
	fmt.Println(u1)
	var u2 uint16 = 65535
	fmt.Println(u2)
	var u3 uint32 = 4294967295
	fmt.Println(u3)
	var u4 uint64 = 18446744073709551615
	fmt.Println(u4)

	var i1 int8 = -128
	fmt.Println(i1)
	var i2 int16 = -32768
	fmt.Println(i2)
	var i3 int32 = -2147483648
	fmt.Println(i3)
	var i4 int64 = -9223372036854775808
	fmt.Println(i4)

	var f1 float32 = 3.14
	fmt.Println(f1)
	var f2 float64 = 3.14159265358979323846
	fmt.Println(f2)
	
	var emptyusername string = ""
	fmt.Println(emptyusername)
}