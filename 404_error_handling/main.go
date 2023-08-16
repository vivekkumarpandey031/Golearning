package main

import (
	//"bytes"
	"fmt"
	//"os"
)

// func GetFileCharCount(Filename *string)(int,error)  {

// if Filename==nil{
// 	return 0, errors.New("nil filename")
// }

// 	bytes ,err := os.ReadFile("data.txt")
// 	if err!=nil{
		
// 		fmt.Println(err)
// 	}
// 	//fmt.Println("count the chars in hte file :", len(bytes))
//    return 
// }
// func main()  {

// 	fileNmae = new(string)
// 	*fileName = "data.txt"
//     fileops.GetFileCharCount(fileName)
// }

// ------->  panic 

func main()  {

	fn1:="victoria"
	ln1:="secret & co"
	fn2:="victoria"
	str1:=GetFullName(&fn1,&ln1)
	str2:=GetFullName(&fn2,nil)
	fmt.Println(*str1,*str2)
}
func GetFullName(firstName,lastName*string) *string {
	if firstName == nil || *firstName == " "{
		panic("firstName cant be nil")
	}
	if lastName == nil || *lastName == " "{
		panic("lastName cant be nil")
	}
	str:= *firstName + *lastName
	return &str
}

  


