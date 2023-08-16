package main

import (
	"fmt"
	//"runtime"
	//"runtime"
	"sync"
	//"math"
	//"math/rand"
)

// main is the main goroutines in Golang programm is it stops working it finish ,
//     the whole program will crash

// func main(){
// 	go fmt.Println("hello victoria secret")
// 	for i:=0;i<=1000;i++{
// 		go IsEven(i)
// 		// num:=rand.Intn(1000)
// 		// go IsEven(num)
// 	}
// 	 runtime.Goexit()

// }

func main() {
	wg := new(sync.WaitGroup)

	i := 1
	wg.Add(101)
	go func() {
		//wg.Add(1)
		for i = 1; i <= 100; i++ {
			 func(num int) {
				//wg.Add(1)
				if i%2 == 0 {
					fmt.Println(num, "is even")
				} else {
					fmt.Println(num, "is odd")
				}
				wg.Done()
			}(i)

		}
		wg.Done()
	}()
	wg.Wait()
	//	runtime.Goexit()

}

// func IsEven(num int)  {

// 	if num%2==0{
// 		fmt.Println(num,"is even")
// 	}else{
// 		fmt.Println(num,"is odd")
// 	}

// }

//  func Fatal(msg string,code int)  {
// 	fmt.Println()
//  }
