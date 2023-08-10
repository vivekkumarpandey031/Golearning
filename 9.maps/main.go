package main

import (
	"fmt"
	"math/rand"
)

 

func main() {

   var mymap map[string]any
   mymap = make(map[string]any)
   mymap["52201"]= "Guntur-1"
   mymap["52206"]= "Guntur-2"
   mymap["52207"]= "Guntur-3"
   mymap["522096"]= "Guntur-4"

   for key,value:=range mymap{
	fmt.Println(key,value)
   }

   val,ok:=mymap["522001"]
   fmt.Println(val,ok)
   delete(mymap,"522001")

   // examples
  
    slice1:=make([]int,100) // create a slice

	for i,_ :=range slice1{
		slice1[i] = rand.Intn(50)
	}

	fmt.Println(slice1)
	mymap1:=make(map[int]uint16)

	for _,v:= range slice1{
		val,ok :=mymap1[v]
		if ok {
			mymap1[v] = val +1;

		}else {
			mymap1[v]+=1;
		}

	}
	fmt.Println(mymap1)
	///create here the taskkk


	//*************************************************

	



}