package main

import (
	//"demo/handlers"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)
var (
	PORT string
)


func check(err error){
	panic(err)
}
func main()  {
PORT = os.Getenv("PORT")
if PORT == ""{
	flag.StringVar(&PORT, "port", "9091", "--port=9091 or -port=9091")
}

//PORT = *flag.String("port","9091","--port=9091 or -port = 9091") 
flag.Parse()
fmt.Println(PORT)
fmt.Println("servweer get started and listening on port",PORT)


	fmt.Println("server started and Listening on port 9091 ")

http.HandleFunc("/ping",func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"pong")
	
})
http.HandleFunc("/greet",greet)
http.HandleFunc("/greetby",GreetBy)
   go http.ListenAndServe("0.0.0.0:9092",nil)
	err:=http.ListenAndServe("0.0.0.0:9091",nil)
	fmt.Println(err,"hello world")
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}

	
}

/*listen&serve
   flag is nothung but to process command line argument
   os.arg vs flag --> you need toi differentiate what is an optional command line argument
   flag provide you all, the facilities to fetch the all datatypes


*/
func greet(w http.ResponseWriter, r *(http.Request))  {
	fmt.Fprintln(w,"hello Victoria Secret & co.")
}

func GreetBy(w http.ResponseWriter, r *(http.Request)){
	if r.Method == "POST"{
	 g:=Greet{}
	 err:=json.NewDecoder(r.Body).Decode(&g)
	if err!=nil{
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprint(w,"Data has been read",g)
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
if err != nil {
    panic(err)
}

defer f.Close()

if _, err = f.WriteString("data has been read by me.  "); err != nil {
    panic(err)
}
}else{
	    w.Write([]byte("Unimplemented method"))
		w.WriteHeader(http.StatusNotImplemented)
}


}
type Greet struct{
	Message string `json:"message"`
	Name string  `json:"name"`

}

func Sayhihttp(w http.ResponseWriter, r *(http.Request)){
	fmt.Fprintln(w,"Hello mrs. victoria ")
}
// type Greet struct{
// 	Message string `json:"message"`
// 	Name string `json:"name"`
// }

type Employee struct{
	No string
	Name string
	Email string
	Mobile string
	Status string
	LastModified int64
}