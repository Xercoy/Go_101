/*
Problem #16: Sine Function
Author: Corey Prak
Date Created: 05/06/2014

Description: Read an integer or floating point value from the command line. It will come into your program as argv[1]. Remember that argv[1]’s type is a pointer to a character that begins a sequence of characters that end in a null byte. In other words, it is what is agreed upon as a string.

For example,
bash>./a.out 4.3

Print out the trigonometric sine value of this number.  You would do this by calling the proper function that is found in Appendix D of you textbook.

Comments:
Observing the math package found at http://golang.org/pkg/math/#Sqrt it can easily be observed that the function Sin does exactly what we need it to do! I <3 Go.

Also was pretty cool having to use the strconv package.
*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// Variable which will hold the user's floating point argument.
	var cmdLineInput float64
	var parseError error

	// Make sure that there's at lease some kind of argument supplied. We can't assume.
	if len(os.Args) < 2 {
		fmt.Printf("\nERROR, requires floating point argument. Exiting...\n\n")
		return
	}

	// The command line argument should be a float, but is of type string. Use ParseFloat to interpret it as a 64 bit floating point number.
	cmdLineInput, parseError = strconv.ParseFloat(os.Args[1], 64)
	if parseError != nil {
		fmt.Printf("\nERROR: %v. Exiting...\n\n", parseError)
		return
	}

	fmt.Printf("\nThe sine of %f is %f.\n\n", cmdLineInput, math.Sin(cmdLineInput))
}
