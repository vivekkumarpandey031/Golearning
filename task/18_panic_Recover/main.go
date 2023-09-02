package main

import (
	"fmt"
	"reflect"
)

func main() {
	Calc("Hi", "Victoria")  
	Calc([]int{10, 20}, 10) 
	sum, sub, mul, modulus := Calc(10, 20)
	fmt.Println(sum, sub, mul, modulus)
}

func Calc(a, b any) (float64, float64, float64, float64) {
	defer RecoverMe()
	var sum, sub, mul, mod float64

	switch a.(type) { 
	case int:
		sum = float64(a.(int) + b.(int))
		mul = float64(a.(int) * b.(int))
		sub = float64(a.(int) - b.(int))
		mod = float64(a.(int) / b.(int))
	case uint:
		sum = float64(a.(uint) + b.(uint))
		mul = float64(a.(uint) * b.(uint))
		sub = float64(a.(uint) - b.(uint))
		mod = float64(a.(uint) + b.(uint))
	case uint8:
		sum = float64(a.(uint8) + b.(uint8))
		mul = float64(a.(uint8) * b.(uint8))
		sub = float64(a.(uint8) - b.(uint8))
		mod = float64(a.(uint8) / b.(uint8))

	case uint16:
		sum = float64(a.(uint16) + b.(uint16))
		mul = float64(a.(uint16) * b.(uint16))
		sub = float64(a.(uint16) - b.(uint16))
		mod = float64(a.(uint16) / b.(uint16))

	case uint32:
		sum = float64(a.(uint32) + b.(uint32))
		mul = float64(a.(uint32) * b.(uint32))
		sub = float64(a.(uint32) - b.(uint32))
		mod = float64(a.(uint32) / b.(uint32))
	case uint64:
		sum = float64(a.(uint64) + b.(uint64))
		mul = float64(a.(uint64) * b.(uint64))
		sub = float64(a.(uint64) - b.(uint64))
		mod = float64(a.(uint64) / b.(uint64))
	case int8:
		sum = float64(a.(int8) + b.(int8))
		mul = float64(a.(int8) * b.(int8))
		sub = float64(a.(int8) - b.(int8))
		mod = float64(a.(int8) / b.(int8))

	case int16:
		sum = float64(a.(int16) + b.(int16))
		mul = float64(a.(int16) * b.(int16))
		sub = float64(a.(int16) - b.(int16))
		mod = float64(a.(int16) / b.(int16))

	case int32:
		sum = float64(a.(int32) + b.(int32))
		mul = float64(a.(int32) * b.(int32))
		sub = float64(a.(int32) - b.(int32))
		mod = float64(a.(int32) / b.(int32))

	case int64:
		sum = float64(a.(int64) + b.(int64))
		mul = float64(a.(int64) + b.(int64))
		sub = float64(a.(int64) - b.(int64))
		mod = float64(a.(int64) / b.(int64))

	case float32:
		sum = float64(a.(float32) + b.(float32))
		mul = float64(a.(float32) * b.(float32))
		sub = float64(a.(float32) - b.(float32))
		mod = float64(a.(float32) / b.(float32))

	case float64:
		sum = a.(float64) + b.(float64)
		mul = a.(float64) * b.(float64)
		sub = a.(float64) - b.(float64)
		mod = a.(float64) / b.(float64)

	default:
		s := fmt.Sprintf("Invalid type detected: %v", reflect.TypeOf(a))
		panic(s)
	}

	return sum, sub, mul, mod

}

func RecoverMe() {
	if pan := recover(); pan != nil {
		fmt.Println(pan, "recovered")
	}

}
