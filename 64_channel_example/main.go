package main

import (
	"fmt"
	"strings"
)


func main()  {


	ch1:=make(chan string)
	ch2:=make(chan string)
	ch3:=MultipleChanOut(ch1,ch2)
	ch11:=make(chan int)
	ch12:=make(chan int)
	ch13:=make(chan int)
	// go func ()  {
	// 	go func ()  {
			
	// 	}
		ch11<-11
		ch12<-12
		ch13<-13
	}()

	go func ()  {
		  ch1 <-"victoria"
          ch2<-"secret & co"

	}()

	for val:=range ch3{
		fmt.Println(val)
	}
	a:= <-VariadicChan(ch11,ch12,ch13)
	fmt.Println(a)

	
}

func MultipleChanOut(ch1,ch2 chan string)  chan string{
	ch:=make(chan string)
     go func()  {
		ch<- <-ch1 +" "+ <-ch2
		close(ch)
	 }()
	return ch
	
}

func SingleToMultiple(ch1 chan string)(ch2 chan string,ch3 chan string)  {
	ch2,ch3=make(chan string),make(chan string)
	
	go func ()  {
		val:=<-ch1
        vals:=strings.Split(val, " ") // ASSUMING THAT there is one space
         ch2 <-vals[0]
		 ch3 <-vals[1]
		 close(ch2)
		 close(ch3)
	}()


	return ch2,ch3
}

func GenrateSquares(num uint)chan int{
	ch:=make(chan int)
	go func ()  {
		for i:=0;i<int(num);i++{
			ch <-i*i
		}
		close(ch)
	}()
	return ch
}

func VariadicChan(chs ...chan int)chan int{
	chx:=make(chan int)
	//kaju_katli:=make(chan int)  -> no need to make this channel this is already one
	go func ()  {
          s:=0    // assining sum 
		for  _,kaju_katli :=range chs{
			s += <-kaju_katli   // add channel->value to sum 
			
		}
			chx<-s
			close(chx)
		
	}()
	return chx
}


func ClosedChannel(ch chan int) {

    for {

        _, ok := <-ch

        if ok {

            select {

            case <-ch:

                fmt.Println("We got the value: ", <-ch)

            }

 

        } else {

            break

        }

    }

}

 


// func FanIn()