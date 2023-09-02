

package main

import "fmt"

func main() {
	
	Check_Ticket("AP", "F", 150, 24)
	Check_Ticket("Delhi", "M", 90, 4)
	Check_Ticket("Karnataka", "F", 150, 24)
	Check_Ticket("Karnataka", "M", 130, 8)
	Check_Ticket("UP", "F", 150, 24)
	Check_Ticket("delhi", "F", 120, 9)

}

func Check_Ticket(State string, Gender string, Height int, Age int) {
	switch {

	case State == "Karnataka" && Gender == "F" || Gender == "M" && Height < 110 && Age < 5:
		fmt.Println("No Ticket")
	case State == "AP" && (Gender == "F" || Gender == "M") && Height < 110 && Age < 5:
		fmt.Println("No Ticket")
	case State == "Delhi" && Gender == "F" || Gender == "M" && Height < 130 && Age < 7:
		fmt.Println("No Ticket")
	case State == "UP" && (Gender == "F" || Gender == "M") && Height < 120 && Age < 6:
		fmt.Println("No Ticket")
	default:
		fmt.Println("Full Ticket")
	
	}

}