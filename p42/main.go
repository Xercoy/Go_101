/*******************************************************************************
Problem #42: Recursive Factorial
Author: Corey Prak
Date Created: 04/03/2014
Description:
Write a program that computes the factorial of a number. Main should read the
value of 0 or a positive integer value from the command line and then output
the factorial of that number. A for loop should be used and the function
should be iterative.
Comments:
*******************************************************************************/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func factorial(number int) int {
	product := 1

	for i := 1; i <= number; i++ {
		product *= i
	}

	return product
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("Must provide an integer greater than or equal to 0.")
	}

	number, _ := strconv.Atoi(os.Args[1])

	fmt.Printf("\nThe factorial of %d is %d.\n\n", number, factorial(number))
}
