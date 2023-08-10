package main

import (
	"fmt"


)

func sum1(a, b, c int) int{
   
   ans:=a+b+c
   return ans
   
}
func main()  {

	var arr = [3][3]int{ { 1, 2, 3}, { 4, 5, 6}, { 7, 8, 9}}
    
	var k int
	for k=0;k<3;k++{
		fmt.Println(sum1(arr[k][0],arr[k][1],arr[k][2]))
  
	   }
	for k=0;k<3;k++{
      fmt.Println(sum1(arr[0][k],arr[1][k],arr[2][k]))

	 
	}
	
	

}