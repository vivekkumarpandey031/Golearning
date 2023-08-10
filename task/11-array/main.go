package main

import (
	"fmt"

	
)

func change(a,b int) []int{
   a,b = b,a
   return []int{a, b}
   
}
func main()  {
     
	var arr = [3][3]int{ { 1, 2, 3}, { 4, 5, 6}, { 7, 8, 9}}
	
	
	/* 1 2 3
	   4 5 6
	   7 8 9 */
	for k:=0;k<3;k++{
		for i:=0;i<3;i++{
			change(arr[i][k],arr[k][i])
		}
	}
	for k:=0;k<3;k++{
		for i:=0;i<3;i++{
			fmt.Print(arr[i][k]," ")
		}
		fmt.Println(" ")
	}
	

}