package main

import (
	"errors"
	"fmt"

)

func main() {
	var mymap map[string]any
	mymap = make(map[string]any)
	mymap["a"] = 1
	mymap["b"] = 2
	mymap["c"] = 6
	mymap["f"] = 4
	fmt.Println("initialy map is this : ",mymap)
	b1, err1 := Delete(mymap, "f")
	fmt.Println("deleting..... ", b1,err1)
	b2, err2 := Update(mymap, "c", 3)
	fmt.Println("updating....", b2, err2)
	fmt.Println("after operation  map is this : ",mymap)

}

func Delete(mymap map[string]any, key string) (bool, error) {
	if mymap == nil {
		return false, errors.New("map is nil")
	}

	_, exists := mymap[key]
	if exists {
		delete(mymap, key)
		return true, nil
	} else {
		return false, nil
	}
}

func Update(mymap map[string]any, key string, value any) (bool, error) {
	if mymap == nil {
		return false, errors.New("map is nil")
	}

	_, exists := mymap[key]
	if !exists {
		return false, errors.New("key does not exist")
	} else {
		mymap[key] = value
		return true, nil
	}
}
