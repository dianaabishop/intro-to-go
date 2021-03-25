package main

import (
	"fmt"
	"strconv"
)

func determineIfCanDrink (person Person) (bool, error) {
	// try to convert age to int
	intAge, err := strconv.Atoi(person.Age)
	if err != nil {
		// if there is an error converting the age value to an int
		// return false (which is the nil value for a bool type)
		// and also return the error 
		return false, err
	}

	// check age
	if intAge >= 21 {
		return true, nil
	} else {
		return false, nil
	}
}

type Person struct {
	Age string
	FirstName string
	LastName string 
}

func main() {
	person1 := Person {
		FirstName: "Diana",
		LastName: "Bishop",
		Age: "32",
	}

	person2 := Person {
		FirstName: "Jason",
		LastName: "Bishop",
		Age: "39",
	}

	people := []Person{person1, person2}
	for _, person := range people {
		fmt.Printf("First Name: %s\n", person.FirstName)
	}
}

