package main

import "fmt"

func main(){
   task5()
}

func task5() {
 // variable declared
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
	// variable value assigned
	v1 = -3
	v2 = 12
	v3 = -3465
	v4 = 1346
	v5 = -6864
	v6 = 4672
	v7 = 345.8930
	v8 = -23937
	v9 = 25382
	v10 =67.364739
	v11 = 34
	v12 = 55
	v13 = 7
	v14 = 'Z'
	//Need to type convert
	o1 := uint8(v1) * v2
	m1 := uint8(v1) - v2
	fmt.Println(o1)
	fmt.Println(m1)
	//Need to type convert
	o2 := uint16(v3) + v4
	m2 := uint16(v3) - v4
	fmt.Println(o2)
	fmt.Println(m2)

	fmt.Println(int(v12))

	//Need to type convert
	o3 := uint32(v5) + v6
	m3 := uint32(v5) - v6
	fmt.Println(o3)
	fmt.Println(m3)
	//Need to type convert
	o4 := int64(v10) + v8
	m4 := int64(v10) - v8
	fmt.Println(o4)
	fmt.Println(m4)

	//Need to type convert
	o5 := uint64(v8) + v9
	m5 := uint64(v8) - v9
	fmt.Println(o5)
	fmt.Println(m5)

	// Need to type convert
	o6 := float64(v7) + v10
	m6 := float64(v7) - v10
	fmt.Println(o6)
	fmt.Println(m6)

	//Need to type convert
	o7 := int(v12) + v11
	m7 := int(v12) - v11
	fmt.Println(o7)
	fmt.Println(m7)

	o8 := int(v13) + v11
	m8 := int(v13) - v11
	fmt.Println(o8)
	fmt.Println(m8)

	o9 := int8(v14) + v1
	m9 := int(int8(v14) - v1)
	fmt.Println(o9)
	fmt.Println(m9)

}