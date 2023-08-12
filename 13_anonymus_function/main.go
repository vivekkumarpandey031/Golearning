package main

import (
	"fmt"
	//"go/constant"
	//"strings"
)

func main()  {
    // normal function call
	Greet()

	// anonymus function syntax and it is called in main function itshelf
	func ()  {
		str:="hello i am farzzziii"
		fmt.Println(str)
	}()

	// another anonymus function call with argument

	func (str string)  {
		fmt.Println(str)

	}("hello baba")

// anonymus fuction fior reversing string
rstr := func (str string)  string{
	reverse := " "
	for _, v := range str{
		reverse = string(v) + reverse
	}
	return reverse
}("hello world")
fmt.Println("Reverse:",rstr)

// dont kniow what is now

f1:=func(str string)string{
	reverse := " "
 for _,v:=range str {
	reverse = string(v)+reverse
 }
 return reverse
}("sai is my beest friend")

fmt.Println(f1)

vo,co,cha:=func (str string)(int,int,int)  {
	vowels:=0
	constants:=0
	chara:=0
   for _,v:= range str{
	 ch:=string(v)
	 if ch!=" " {
		if v>=65 && v<=90 || v>=97&&v<=122{
            if ch=="a"||ch=="e"||ch=="i"||ch=="o"||ch=="u" || ch=="A"||ch=="E"||ch=="I"||ch=="O"||ch=="U"{
				vowels++
			}else{
				constants++
			}
		} else{
			chara++
		}

	 }
   }
   return vowels,constants,chara


	
}("Victoriahghdg#$%$$%^%&*^ Secret @ and co.")
fmt.Println(vo,co,cha)
 
/* create a map with 
      key --> add               value--> value by anonimus function
	          sub
              mulktilply
			division

			*/
		 
			
var mp map[any]func(int,int)int

mp = make(map[any]func(int, int) int)
mp["add"] = func (a,b int)int  {
	an:=a+b
	return an
}
mp["sub"] = func (a,b int)int  {
	an:=a-b
	return an
}
mp["multiply"] = func (a,b int)int  {
	an:=a*b
	return an
}
mp["division"] = func (a,b int)int  {
	an:=a/b
	return an
}
for key,fn:=range mp{
	fmt.Println(key,fn(4,2))

}










	
}
//--> what is an anony us function
//-->write an anonymus function that should return number of vowels,
//    consonants and special characters in a given strings.


func Greet(){
	fmt.Println("hello welcome to disney land")
}