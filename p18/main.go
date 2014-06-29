/*
Problem #18: Solid Box of Asterisks
Author: Corey Prak
Date Created: 04/03/2014

Description: Your task is to write a C program that accepts two positive integer values from the keyboard, say L and H. Both L and H will be less than 21. Your program then prints out a solid of box of asterisks, with L horizontal stars and H vertical stars. The box is filled with asterisks. For example,

If you were to type in 5 and 3, you output would be:

*****
*****
*****

Comments: YAY, we're finally at loops!!!
*/

package main

import (
	"fmt"
)

func main() {

	//Integer variables that will contain user submitted length and height.
	var length int
	var height int
	var scanErr error

	//Integer values that will act as incrementers. Starting with i by convention. No particular reason for multiple declarations in one line, just wanted to show that it was possible.
	var i, j int

	fmt.Printf("\n\nPlease enter two integer values which will be used to determine the size of a box of asterisks. Both values must be less than 21, for length and height, respectively: ")

	//Also possible to read in one integer at a time, just wanted to show that this was possible as well. In my limited experience with golang's format specifiers, they must not have any spaces between them; compiler will complain.
	_, scanErr = fmt.Scanf("%d %d", &length, &height)
	if scanErr != nil {
		fmt.Printf("\n\nERROR: %v\n\nExiting...\n\n", scanErr)

		return
	}

	//Debugging
	fmt.Printf("\nLength:%d, Height:%d\n", length, height)

	//There are so many ways to design this portion of error checking. It's possible to do some error handling, or have a loop that will continuously keep asking for valid input, or something entirely different...
	if (length >= 21) || (height >= 21) {
		fmt.Printf("Invalid input for length(%d) and/or height(%d). Exiting...\n\n", length, height)
		return
	}

	//Spacing.
	fmt.Printf("\n\n")

	//i is already set to its default value of 0 when declared, just being explicit to demonstrate the first format of a for loop.
	//Demonstrates the first format of a loop: for *init incrementer* ; *condition* ; *incrementer modification* {}
	//Increment modification simply adds one to the current value of i, assigns back to the variable.
	for i = 0; i < height; i = i + 1 {

		//Demonstrates the second form of a for loop, albeit simpler: for *condition* {}
		//Note the particular notation that modifies the increment. j += 1 is synonymous with j = j + 1.
		for j < length {
			fmt.Printf("*")

			j += 1
		}

		fmt.Printf("\n")
		j = 0
	}

	//Spacing.
	fmt.Printf("\n\n")
}
