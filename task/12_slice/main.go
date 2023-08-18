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
		// second largest numbers
		if slice1[i] > ln{
			sln = ln
			ln = slice1[i]

		}else if slice1[i] > sln{
			sln = slice1[i]
		}
	// second smallest number
		if slice1[len(slice1)-i-1] < sn {
			ssn =sn
			sn = slice1[len(slice1)-i-1]
			

		}else if slice1[len(slice1)-i-1] < ssn{
			ssn=slice1[len(slice1)-i-1]
			
		}
		//fmt.Println(sln,"     ",ln)
	}
	return sln, ssn

}
func main() {

	var arr = []int{2,1,3,6,5}
	slice1 := arr[:]
	fmt.Println(slice1)
	sln, ssn := getBig_getSmall(slice1)
	fmt.Println("Second greatest number in the slice is :", sln, "\n", "Second smallest number in the slice is :", ssn)

}
