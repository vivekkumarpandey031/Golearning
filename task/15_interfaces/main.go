package main

import (
	"fmt"
	"math"
)

type geometry interface{
	area()  float64
	perim() float64
}

type circle struct{
	raidus float64
}

func (c circle) area()float64  {
	return math.Pi*c.raidus*c.raidus
}

func (c circle) perim() float64{
	return 2*math.Pi*c.raidus
}

func main(){
	c1:=circle{raidus: 3}

	fmt.Println(c1.area())
	fmt.Println(c1.perim())

}