package main

import "log"

// type User struct {
// 	FirstName string
// 	LastName  string
// }

func main() {

	// myMap := make(map[string]User) // Cannot assume a map is sorted // always use this sintax to create a map

	// me := User{
	// 	FirstName: "Yves",
	// 	LastName:  "ASdasdnasd",
	// }

	// myMap["me"] = me

	// // myMap["dog"] = "Samson"
	// // myMap["other-dog"] = "Cassie"

	// // myMap["dog"] = "Fido"

	// // log.Println(myMap["dog"])

	// // myMap["First"] = 1
	// // myMap["Second"] = 2

	// // log.Println(myMap["First"], myMap["Second"])
	// log.Println(myMap["me"].FirstName)

	// var myNewVar float32

	// myNewVar = 11.1

	// var mySlice []int

	// mySlice = append(mySlice, 1)
	// mySlice = append(mySlice, 3)
	// mySlice = append(mySlice, 2)

	// sort.Ints(mySlice)

	// log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[:3])

	names := []string{"one", "seven", "fish"}

	log.Println(names[1:3])

}
