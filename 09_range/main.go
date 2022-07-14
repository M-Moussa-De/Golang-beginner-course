package main

import "fmt"

func main() {

	ids := []int{15, 58, 2, 24, 35, 87}
 
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
     sum += id
	}
	fmt.Printf("IDs sum is %d\n", sum)

	
	// Range with map
	mp := map[string]string{"job": "Web developer", "skills": "HTML, CSS"}
	
	for k, v := range mp {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

}