package main

import "fmt"

func main() {

	// Define map
	emails := make(map[string] string)

	emails["John"] = "john@gmail.com"
	emails["Doe"] = "doe@gmail.com"

	fmt.Println("Length ", len(emails)) // 2
	
	fmt.Println("Map ", emails)

	fmt.Println(emails["John"])

	emails["test"] = "test@email.com"

	fmt.Println("Before deleting ", emails)

	// Delete from map
	delete(emails, "test")

	fmt.Println("After deleting ", emails)

	// Declare map and add key value
	addresses := map[string]string{"Mark": "Str. 10, 12345"}
	addresses["Alessandro"] = "Str. 12, 18975"

	fmt.Println(addresses)


}