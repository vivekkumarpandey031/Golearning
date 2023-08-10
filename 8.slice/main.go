package main

import "fmt"

func main()  {
	// slice1:=[]int{10,20,30}
	// arr1:=[5]int{10,11,12,13,14}
	// slice2:=arr1[:]
	// slice3:=arr1[:3]
	// slice4:=arr1[3:]
    // slice5:=arr1[1:4]
	// fmt.Println(slice1,slice2,slice3,slice4,slice5,arr1)
	// slice6:=make([]int, 4)
	// slice7 := make([]int,3,6)
    // fmt.Println(slice6,slice7)
	// slice8:=make([]int, 2)
	// slice8[0]=87
	// slice8[1]=92
	// slice9:=slice8
	// slice9[0]=97
	// slice9[1]=102
	// fmt.Println(slice8,slice9)
	// slice8 = append(slice8, 103)
	// fmt.Println(slice8,slice9)
	// slice10:=make([]int, len(slice8))
	// copy(slice10,slice8)
	// slice10[0]=207
	// fmt.Println(slice10,slice8

	s:=[]int{1,2,3,4,5}
	r:=s
	r = append(s[:2],s[3:]...)
	fmt.Println(r,s)
	// arr := [...]int{1, 2, 3, 4, 5}
// fmt.Println(arr) // [1 2 3 4 5]
// sl := arr[2:4]
// fmt.Println(sl) // [3 4]
// sl[0] = 0
// fmt.Println(sl) // [0 4]
// fmt.Println(arr)
// sl = append(sl, 9)
// fmt.Println(sl) // [0 4 9]
// fmt.Println(arr) // [1 2 0 4 9]



	


}