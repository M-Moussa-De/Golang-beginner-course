package main

import "fmt"

func main() {

	a := 100
	b := &a
	c := *b

	fmt.Println(a, b, c) // 100 0xc000bbb015 100

	*b = 200

	fmt.Println(a, b, c) // 200 0xc000bbb015 100 


	fmt.Printf("%T\n", b) // *int

	// Read value from pointer to
	fmt.Println(*b) // 200
	fmt.Println(*&a) // 200
}