/*
Problem #30: Unfilled Box
Author: Corey Prak
Date Created: 05/14/2014

Description: Read a pair of positive integers, L and H from standard input, which will be less than 21, creating a box of asterisks.

For example, if the input data is 6 and 4 then the output should be:
******
*    *
*    *
******

Comments:
*/

package main

import (
	"fmt"
)

func main() {
	var length, height int
	var err error
	var i, j int

	fmt.Printf("\nPlease enter two integer values for length and height, respectively: \n")

	_, err = fmt.Scanf("%d %d", &length, &height)
	if err != nil {
		fmt.Printf("Scanf Error: %v\n\nExiting...\n\n", err)
		return
	}

	if (length >= 21) || (height >= 21) {
		fmt.Printf("\nInput Error: Invalid input. Length and height must be less than 21.")
		return
	}

	fmt.Printf("\n")

	for ; i < height; i++ {
		for ; j < length; j++ {

			if (i == 0) || (i == height-1) {
				fmt.Printf("*")
			} else if (j == 0) || (j == length-1) {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Printf("\n\n")

		j = 0
	}
}
