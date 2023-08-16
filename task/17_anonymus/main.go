package main

import "fmt"

func main()  {
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
}