package main

import (
	"fmt"
	"math/rand"
)

 

func main() {

 
	var p1 Person //declaration
	p1.Id = 


	



}
type Person struct{
	Id    int
	Name  string
	Email string
	Mobile string
	addr   Address   // composition ?? embbeddeed struct
}
type Address struct{
    Line1, City, Street, State, Country, PinCode  string 
}