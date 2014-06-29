/*******************************************************************************
Problem #33: Persistence of a Number
Author: Corey Prak
Date Created: 06/29/2014
Comments:
*******************************************************************************/
package main

import (
	"fmt"
	"io"
	"log"
)

func persistence(number int) int {
	product, persistenceNumber := 1, number

	for persistenceNumber > 10 {
		for temp := persistenceNumber; temp > 0; temp = temp / 10 {
			product = product * (temp % 10)
		}

		persistenceNumber = product
		product = 1
	}

	return persistenceNumber
}

func main() {
	var userInput int

	for {
		fmt.Printf("\nPlease enter a positive integer: ")
		_, scanErr := fmt.Scanf("%d", &userInput)
		if scanErr == io.EOF {
			log.Fatalln("EOF entered. Program exiting...")
		} else if userInput <= 0 {
			log.Fatalf("ERROR: Positive integer Needed\n\n")
		}

		fmt.Printf("The persistence number of %d is %d\n", userInput, persistence(userInput))
	}

	fmt.Printf("\n\n")
}
