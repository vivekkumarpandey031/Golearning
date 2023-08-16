package main

import (
	"fmt"
	//"strings"
)


func main()  {

ch:=make(chan int)
ch <-100
ch <-200
fmt.Println(<-ch, <-ch)
ch<-300
fmt.Println(<-ch,<-ch)
	
	
}

// func MultipleChanOut(ch1,ch2 chan string)  chan string{
// 	ch:=make(chan string)
//      go func()  {
// 		ch<- <-ch1 +" "+ <-ch2
// 		close(ch)
// 	 }()
// 	return ch
	
// }

// func SingleToMultiple(ch1 chan string)(ch2 chan string,ch3 chan string)  {
// 	ch2,ch3=make(chan string),make(chan string)
	
// 	go func ()  {
// 		val:=<-ch1
//         vals:=strings.Split(val, " ") // ASSUMING THAT there is one space
//          ch2 <-vals[0]
// 		 ch3 <-vals[1]
// 		 close(ch2)
// 		 close(ch3)
// 	}()


// 	return ch2,ch3
// }

// func GenrateSquares(num uint)chan int{
// 	ch:=make(chan int)
// 	go func ()  {
// 		for i:=0;i<int(num);i++{
// 			ch <-i*i
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }

// func VariadicChan(chs ...chan int)chan int{
// 	chx:=make(chan int)
// 	//kaju_katli:=make(chan int)  -> no need to make this channel this is already one
// 	go func ()  {
//           s:=0    // assining sum 
// 		for  _,kaju_katli :=range chs{
// 			s += <-kaju_katli   // add channel->value to sum 
			
// 		}
// 			chx<-s
// 			close(chx)
		
// 	}()
// 	return chx
// }


// func ClosedChannel(ch chan int) {

//     for {

//         _, ok := <-ch

//         if ok {

//             select {

//             case <-ch:

//                 fmt.Println("We got the value: ", <-ch)

//             }

 

//         } else {

//             break

//         }

//     }

//}

 


// func FanIn()