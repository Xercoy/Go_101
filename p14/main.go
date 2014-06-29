/*
Problem #14: Argc
Author: Corey Prak
Date Created: 04/20-something/2014
Date Completed: 04/28/2014

Description:
Write a C program that prints the number of command line arguments. This is the number that is stored in the argument: argc.  Do not lose points because your code has a ragged indent or a misspelled word. For example,

bash>  ./a.out would print out 1
bash>  ./a.out 511 Mel Ott Life is good               would print out 7
bash>  ./a.out  Jem and Scout                                     would  print out 4


Comments:
//	fmt.Printf("\nThere are %d command line arguments.\n\n", strconv.atoi(os.Args()));
^lol, why the heck did I write this?
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	var argLength int = len(os.Args)

	if argLength != 1 {
		fmt.Printf("\nThere are %d command line arguments.\n\n", argLength)
	} else {
		fmt.Printf("\nThere is one command line argument.\n\n", argLength)
	}
}
