package main

import "fmt"

func main() {
	// Arrays => must have a fest length and can't be expanded once it's declared
	// arr :=  [3]string{"Banana", "Orange", "Pineapple"}

	// Slice => In opposite of arrays, its size is dynamic
	arraySlice := []string{"Banana", "Orange", "Pineapple"}

	fmt.Println(arraySlice[0:3]) // Range prints the elements at index from 0 to 2 but 3 not included

	fmt.Println(len(arraySlice))  // Length

	arraySlice = append(arraySlice, "Mango") // Append one new element
	arraySlice = append(arraySlice, "Peach", "Apple", "Stawberry") // Append several elements

	arraySlice[0] = "Kiwi" // Replace element with another

	fmt.Println(arraySlice)
}