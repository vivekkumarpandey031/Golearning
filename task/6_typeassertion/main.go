package main

import "fmt"

func main() {

	var (
		v1  int8
		v2  uint8
		v3  int16
		v4  uint16
		v5  int32
		v6  uint32
		v7  float32
		v8  int64
		v9  uint64
		v10 float64
		v11 int
		v12 byte
		v13 uint
		v14 rune
	)
	v1 = -4
	v2 = 123
	v3 = -1234
	v4 = 12345
	v5 = -56789
	v6 = 34567
	v7 = 2345.1234
	v8 = -2383937
	v9 = 2536382
	v10 = 25367.364739
	v11 = 234
	v12 = 255
	v13 = 47
	v14 = 'V'

	var iface1, iface2 any
	iface1 = 2
	iface2 = 3.

	add := int(v1) + int(v2) + int(v3) + int(v4) + int(v5) + int(v6) + int(v8) + int(v9) + v11 + int(v12) + int(v13) + int(v14) + iface1.(int)
	mul := float64(v7) + float64(v9) + float64(v10) + iface2.(float64)

	fmt.Println("Addition:", add)
	fmt.Println("Multiplication", mul)
	fmt.Printf("Type of iface1: %T\n", iface1)
	fmt.Printf("Type of iface2:  %T\n",iface2)

}