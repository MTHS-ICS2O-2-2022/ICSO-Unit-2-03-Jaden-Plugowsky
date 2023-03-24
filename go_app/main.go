// Created by: Jaden Plugowsky
// Created on: March 2023
//
// This program puts together the user's input, in this case, an address

package main

import "fmt"

func main() {
	// This function will put together the inputted address
	var streetNumber int
	var streetName string

	// Input
	fmt.Println("This program puts together a user's address.")
	fmt.Print("\nEnter your house's number: ")
	fmt.Scanln(&streetNumber)
	fmt.Print("Enter your street's name: ")
	fmt.Scanln(&streetName)

	// Output
	fmt.Println("Your address is: ", streetNumber, streetName)

	fmt.Println("\nDone.")
}
