package main

import "fmt"

var num1 = 10
var num2 int = 10

var float1 = 10.45
var float2 float32 = 10.45

var positiveNum uint8 = 100

var bool1 = true
var bool2 bool = false

var str = "This is a string"
var str2 string = "This is a string"

func main() {
  
	// declare variables implicitly

	num3 := 105

    str3 := "This is a string, its type is decleard implicitly"

	fmt.Println(num3)
	fmt.Println(str3)

}