package main

import "fmt"

func Incr(num *int){
	*num++
}

func main(){
	fmt.Println("pointers")
	num1:= 100
	p := &num1 
	fmt.Println(*p)
     

	// ptr4 = new(int)
    

}