/*
Problem #27: Reverse
Author: Corey Prak
Date Created: 05/13/2014

Description: Read in 10 integer numbers from the keyboard, store them in an array, and print them in reverse order.

Comments:
*/

package main

import (
	"fmt"
)

func main() {
	var numberArray [10]int
	var userInput, i int
	var err error

	fmt.Printf("\nPlease enter ten integer values:\n\n")

	for ; i < 10; i++ {
		_, err = fmt.Scanf("%d", &userInput)
		if err != nil {
			fmt.Printf("\nScan Error: %v\n\nExiting...\n\n", err)
			return
		}

		numberArray[i] = userInput
	}

	fmt.Printf("\nThe 10 integer values displayed in reverse:\n")

	for i = i - 1; i >= 0; i-- {
		fmt.Printf("\n%d", numberArray[i])
	}

	//format
	fmt.Printf("\n\n")
}
