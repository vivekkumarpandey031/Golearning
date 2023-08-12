package main

import (
	"fmt"

	//"golang.org/x/tools/go/analysis/passes/defers"
)
func add(a,b int) any {
	return a+b
}

func main(){
	a:=10
	b:=20
	defer fmt.Println(a,b,add(a,b))


	defer func ()  {
		c:=add(a,b)
		fmt.Println(a,b,c)
	}()
	a=11
	b=21
	c:=add(a,b);
	fmt.Println(a,b,c)
}