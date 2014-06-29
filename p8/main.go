/*******************************************************************************
Problem #8: One Horizontal Line of Asterisks
Author: Corey Prak
Date Created: 04/03/2014

Description:
Create a file called testdata8 that has a single integer value stored into it.
The integer will be a number between 1 and 30 inclusive. Create a C program
that uses the fscanf function to read in this integer value and then prints a
horizontal line containing that many asterisks’ (*). Use a for-loop in your
program.

Comments:

06/04/2014

- Added comments.

- Made HUGE changes to code.

- Decided not to remove short declarations to show how elegant the code could
  be.

- Must REMEMBER TO CLOSE THE FILE. I can't seem to get that into my head.
*******************************************************************************/

package main

import (
	"fmt"
	"os"
)

func main() {

	/* First variable should be self explanatory, second is the interator for i. */
	var numberFromFile, i int

	// Open the file, return address to filePtr, fErr if possible.
	filePtr, fErr := os.Open("testdata8")
	if fErr != nil {
		fmt.Println(fErr)
		return
	}

	/* We're not worried sbout the first return value of the Fscanf function
	   just yet, so omit it with an underscore. */
	_, scanErr := fmt.Fscanf(filePtr, "%d", &numberFromFile)
	if scanErr != nil {
		fmt.Println(scanErr)
		return
	}

	for i = 0; i < numberFromFile; i++ {
		fmt.Printf("*")
	}

	fmt.Printf("\n")

	filePtr.Close()
}
