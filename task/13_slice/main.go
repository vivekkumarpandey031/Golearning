package main

import (
	"fmt"
)

func getBig_getSmall(slice1 []int) (int, int) {
	ln := slice1[0]
	sn := slice1[0]
	for _, i := range slice1 {
		if slice1[i-1] > ln {
			ln = slice1[i-1]
		}
		if slice1[i-1] < sn {
			sn = slice1[i-1]
		}
	}
	return ln, sn

}
func main() {

	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8,9}
	slice1 := arr[:]
	fmt.Println(slice1)
	ln, sn := getBig_getSmall(slice1)
	fmt.Println("greatest number in the aray is :", ln, "\n", "smallest number in the array is :", sn)

}
