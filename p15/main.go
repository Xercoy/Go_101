/*
Problem #15: Using the sqrt Function
Author: Corey Prak
Date Created: 05/06/2014

Description: Write a C program that asks the user to enter a floating point number from the keyboard and then prints out the square root of that number. You should call the sqrt function. You can browse Appendix D to which header file you should include. The textbook is your best friend. Read it. You should know that the sqrt function definition is not located in libc.a. It is located in libm.a and thus you will need to use the –lm option when you try to build the executable image via the gcc utility. You also need to include math.h.

Comments:
The key to solving this problem was simply finding the sqrt function in a library that would be supplied by some package. Google is everyone's best friend.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	//I like forward declarations...
	var userNumber float64
	var readError error

	fmt.Printf("\nPlease enter a floating point number whose square root will be calculated: ")
	_, readError = fmt.Scanf("%f", &userNumber)
	if readError != nil {
		fmt.Printf("\nERROR: %v. Exiting...\n\n")
		return
	}

	fmt.Printf("\nThe square root of %f is %f.\n\n", userNumber, math.Sqrt(userNumber))
}
