package slice1

import (
	"fmt"
	//"math"
)
// func change(a,b int) []int{
// 	a,b = b,a
	
// 	return []int{a, b}
	
//  }

func getSum(slice1 []int)(int)  {
	sum:=0
	for i:=0;i<len(slice1);i++{
	    sum+=slice1[i]
	}
	fmt.Println(sum)
    return sum
		
	
}

func getReverse(slice1 []int)(any)  {
	for i,j:=0,len(slice1)-1;i<j;i,j=i+1,j-1{
		slice1[i],slice1[j]=slice1[j],slice1[i]
	}
	fmt.Println(slice1)
	return slice1
	
}

func getBig(slice1 []int) (int) {

	ln:=slice1[0]
 
 for i:=0;i<len(slice1);i++{
    if slice1[i]>ln{
		ln = slice1[i]
	}
	
 }
 fmt.Println(ln)
 return ln
}
	



func get_small(slice1 []int)(int){
	
 sn:=slice1[0]
 for i:=0;i<len(slice1);i++{
    
	if slice1[i]<sn{
		sn = slice1[i]
	}
 }
 fmt.Println(sn)
 return sn
}
// func main() {

// 	var arr = []int{2,1,3,6,5}
// 	slice1 := arr[:]
// 	fmt.Println(slice1)
// 	sln := getBig(slice1)
// 	ssn := get_small(slice1)
// 	sum:=getSum(slice1)
// 	rev:=getReverse(slice1)
//  	fmt.Println(sln,"\n",ssn,"\n",rev,"\n",sum)

// }
