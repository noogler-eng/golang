package main

import "fmt"

// new datatype created by us
type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	// person1 datatype of struct person
	var person1 Person
	var person2 Person

	person1.firstName = "sharad"
	person1.lastName = "poddar"
	person1.age = 20

	fmt.Println(person1)
	// here firstName and lastName is string so initilized with ""
	// here age is initilized with 0 as int
	fmt.Println(person2)
	// this last will print to int to string '\x00'
	fmt.Printf("%q\n", person2)

	person3 := Person{
		firstName: "anyone",
		lastName:  "something",
		age:       12,
	}
	fmt.Printf("%q\n", person3.firstName)
	fmt.Printf("%q\n", person3.lastName)
	fmt.Printf("%v\n", person3.age)

	// new keyword
	person4 := new(Person)
	person4.firstName = "sha"
	person4.lastName = "rad"
	person4.age = 15
	fmt.Println(person4.age)

}
