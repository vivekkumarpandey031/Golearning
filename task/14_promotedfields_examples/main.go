package main

import "fmt"

func main() {

	//check if promioted field address can be direcly accesible
	var company Company = Company{
		Id:        1,
		name:      "Victoria Secrets & Co",
		website:   "https://www.victoriassecret.com/in/",
		telephone: "9876543210",
		fax:       "987654321000",
		Address: Address{ 
			city:    "Bengaluru",
			line1:   "Manyata embassy tech park",
			street:  "Nagwara",
			state:   "Karnataka",
			country: "India",
			pincode: 56008,
			Map: Map{ 
				lat: 11.11,
				lan: 11.11,
			},
		},
	}

	fmt.Println("Address is directly accesible")
	fmt.Println(company.city,",",company.line1,",",company.street,",",company.country)
	fmt.Println("Map is also directly accesible so every promoted fields will be directly accessible")
	fmt.Println(company.lan,",",company.lat)

}

type Company struct {
	Id        int
	name      string
	website   string
	telephone string
	fax       string
	Address   //promoted field
}

type Address struct {
	city    string
	line1   string
	street  string
	state   string
	country string
	pincode int
	Map     //promoted field
}

type Map struct {
	lat, lan float64
}