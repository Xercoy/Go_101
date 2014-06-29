/*******************************************************************************
Problem #10: Sum of Twenty
Author: Corey Prak
Date Created: 04/17/2014

Description: Problem 10: Sum of Twenty

Write a program that reads in 20 positive integer values from a file called
testdata10 and computes their sum, then prints the value of the sum onto the
screen.  Your program should not use 20 distinct variables. Your program should
not use an array.

Comments:

06/08/2014

- Comments added
- Forgot to close the file. Fixed it.
- I don't dislike short declarations as much as I used to.
- Renamed Fscanf variable countErr to scanErr, horribly named.
- Short declarations

Short declarations are convenient and effective when used as such. I disliked
them when I started learning because it made me lazy and prevented me from
really learning about the types of data that are being declared. In an earlier
version of this assignment I mentioned, "...This was the first time that I
was able to actually declare a file pointer normally, quite frustrating..."

After getting the hang of things, I realized that short declarations could help
organization and readability. One example: suppose you're working with a HUGE
file  and all the declarations are listed at the top before any executable
statements. To have a grasp of the context of the declared variables at any
point, we'd have to constantly refer to the upper portion of the code where
the declarations took place.

Well, if there was a variable of a particular type that was utilized briefly
within an arbitrary part of the code, then there's no harm done. A variable
of type error could be declared with := and manipulated as desired; it only
"exists" within the context of a few lines, and was probably used and declared
within one line.
*******************************************************************************/

package main

import (
	"fmt"
	"os"
)

func main() {
	/* filePointer could be be better named, maybe to file.
	   I personally like to just call it like it is, at least for
	   now. */
	var filePointer *os.File
	var fileNumber, sum int
	var openErr, scanErr error
	var i int

	filePointer, openErr = os.Open("testdata10")
	if openErr != nil {
		fmt.Printf("ERROR: %v\n\nExiting...\n\n", openErr)
	}

	/*  Have the for loop iterate twenty times, reading from the
	    file and adding in the scanned value to the sum variable.
	    Straightforward since we know how many values we want to
	    read in.

	    If we wanted to read an arbitrary number of integers to
	    the file, we could simply keep reading until the scanErr
	    reached the end of the file, specified by making the error
	    variable a value of a constant, EOF, found in the os
	    package.

	    Going back to the code, observe that I omitted the first
	    portion of this type of for loop, which is setting the
	    value of the iterator. i is already at its defaulted
	    value of 0. Convenient! */
	for ; i < 20; i++ {

		/* The underscore (_) is used to omit the first return value
		   which is the number of successful reads. It's omitted
		   since this return value isn't important to our program. */
		_, scanErr = fmt.Fscanf(filePointer, "%d", &fileNumber)

		if scanErr != nil {
			fmt.Printf("ERROR: %v\n\nExiting...\n\n", scanErr)
			return
		}

		/* Same as sum = sum + fileNumber */
		sum += fileNumber
	}

	fmt.Printf("\nThe sum of the 20 integers read from testdata10 is %d\n\n", sum)

	filePointer.Close()
}
