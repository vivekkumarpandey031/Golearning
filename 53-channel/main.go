package main

import "fmt"

// func main()  {
// 	ch:=make( chan int) //Unbuffered channel has been instantiated
// 	ch<-100 // Assign a value to the channel
// 	num:= <-ch  //recieve a value from the channel
// 	fmt.Println(num)
// }

func main()  {
	ch:=make( chan int) //Unbuffered channel has been instantiated
	//ch:=make( chan int,(length_of_the_channel))  --> this is a buffered channel
	go func ()  {
		ch<-100 // Assign a value to the channel
	}()
	num:= <-ch  //recieve a value from the channel
	fmt.Println(num)
}