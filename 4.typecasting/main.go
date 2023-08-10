package main

import "fmt"
func main()  {
	

	var num1 uint8 = 38
	var num2 int = int(num1)
	fmt.Println(num2)
	var num3 int = 65001
	var num4 uint8 = uint8(num3)
	fmt.Println(num4)
	var num6 int = 65
	var str2 string = string(num6)
	fmt.Println(str2)
	//sprint --> does what that it print what argument it tskes as it is
	   // spaces are inserted only if niether of them are string or else they will concatenate.
	str3:=fmt.Sprint(65)
	fmt.Println(str3)
}