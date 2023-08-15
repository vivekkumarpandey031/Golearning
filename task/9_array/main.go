package main

import (
	"fmt"

	
)

func getBig_getSmall(arr[]int)(int,int){
 ln:=arr[0]
 sn:=arr[0]
 for i:=0;i<len(arr);i++{
    if arr[i]>ln{
		ln = arr[i]
	}
	if arr[i]<sn{
		sn = arr[i]
	}
 }
 return ln,sn

}
func main()  {
     
	var arr = []int { 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ln,sn:=getBig_getSmall(arr)
	fmt.Println("greatest number in the aray is :",ln,"\n","smallest number in the array is :",sn)
	
	
	
	

}