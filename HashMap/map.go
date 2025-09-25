package main

import "fmt"

func main() {
	// create a hash map key: string, value: int
	ages := make(map[string]int)

	// insert
	ages["Shiva"] = 19
	ages["Ravi"] = 20
	ages["prince"] = 22

	// access the value as key
	fmt.Println("Shiva's age: ", ages["Shiva"])

	//Update a value
	ages["Shiva"] = 20
	fmt.Println("Shiva's updated age : ", ages["Shiva"])

	// check if a key exists
	if age, ok := ages["Ravi"]; ok {
		fmt.Println("Ravi's age: ", age)
	} else {
		fmt.Println("Ravi not found")
	}

	// delelte a key value pair
	delete(ages, "Anita")

	// loop over all keys and value
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
