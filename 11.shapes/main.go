package main

import (
	"fmt"

)

 

func main() {

 

	fmt.Println("checkied it printed")



}

type Rect struct{
	L,B float32
}
type Square struct{
	S float32
}
func (r Rect)Area() float32 {
   return r.L*r.B
}