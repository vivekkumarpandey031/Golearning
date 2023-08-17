package main

import (
	"fmt"
	"math"
)

func getBig_getSmall(slice1 []int) (int, int) {
	sn := math.MaxInt
	ln:=math.MinInt
	//fmt.Println(sn)
	
	sln := ln+1
	ssn := sn-1
	//fmt.Println(sn)
	 for i:=0;i<len(slice1);i++ {
		//fmt.Println(sn)
		if slice1[i] > ln {

			sln=ln
			ln = slice1[i]
			//fmt.Println(ln," ",sln)
		}
		//fmt.Println(sn)
		if slice1[len(slice1)-i-1] < sn {
			ssn =sn
			sn = slice1[len(slice1)-i-1]
			
		}
		//fmt.Println(slice1[len(slice1)-i-1])
	}
	return sln, ssn

}
func main() {

	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := arr[:]
	fmt.Println(slice1)
	sln, ssn := getBig_getSmall(slice1)
	fmt.Println("Second greatest number in the slice is :", sln, "\n", "Second smallest number in the slice is :", ssn)

}
