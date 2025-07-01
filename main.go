package main

import (
	"log"
	"time"
)

// var s = "seven"

// to make a function or a variable puclic available outside of the package use Capital letter

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	// var s2 = "six"

	// s := "eight" // local variables have priority over global declarations

	// log.Println("s is", s)
	// log.Println("s2 is", s2)

	// saySometring("xxx")

	user := User{
		FirstName:   "Yves",
		LastName:    "Martins",
		PhoneNumber: "(11)91111-1111",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)

}

// func saySometring(s3 string) (string, string) {
// 	log.Println("s from the saySometring func is", s)
// 	return s3, "World"
// }
