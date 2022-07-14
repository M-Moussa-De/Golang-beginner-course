package main

import "fmt"

func main() {

	// if else if  else
	x := 8
	y := 10

	if x < y {
		fmt.Println("x is less than y")
	} else if x > y {
		fmt.Println("x is bigger than y")
	} else {
		fmt.Println("x is same as y")
	}


	// switch
    color := "blue"
	
	switch color {
		case "green":
			fmt.Println("Color is green")
		case "yellow":
			fmt.Println("Color is yellow")
		case "blue":
			fmt.Println("Color is blue")
        default:
			fmt.Println("Color is unknown")
	}
}