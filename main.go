package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySometring()

	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySometring() (string, string) {
	return "sometring", "else"
}
