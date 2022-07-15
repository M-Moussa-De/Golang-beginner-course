package main

import (
	"fmt"
	"strconv"
)

// Define struct
type Person struct {
	fname, lname, city, gender  string
	age                         int
}
// Methods are two kinds => 
                       // 1- value reciever 
					   // 2- pointer reciever

// 1- Value reciever method => Doesn't change the data but only calls/uses it
func (p Person) getInfo() string {
  return "Hello, my name is " + p.fname + " " + p.lname + " and I'm " + strconv.Itoa(p.age) + "."
}

// 2- Pointer reciever method => Changes/Updates the data
func (p *Person) setAge(age int) {
  p.age = age
}

func (p *Person) changeWifeLastnameIfMarried(spouseLastname string) {
  if p.gender == "m" {
	return
  }

  p.lname = spouseLastname
}


func main() {

	// Initialize person using struct
	person1 := Person{fname: "John", lname: "Doe", city: "Boston", gender: "m", age: 45}

	// Alternative
	person2 := Person{"Samantha", "Smith", "New Yourk",  "f",  41}

	fmt.Println(person1, person2)

	fmt.Println(person1.fname) // John

	person1.fname = "Mark"

	fmt.Println(person1.fname) // Mark

	// Call method on a instance
	info := person1.getInfo();
    fmt.Println(info)

	person2.setAge(40) // Was 45, became 40
    fmt.Println(person2.age)

    person2.changeWifeLastnameIfMarried(person1.lname) // Was Möller, became Doe
	fmt.Println(person2.lname)

}