/*
Problem #21: Argv
Author: Corey Prak
Date Created: 04/03/2014

Description: Write a C program that will print out each command line argument on a separate line. Use a for-loop. The loop should have an index variable i that ranges from 0 to (argc-1).  Remember the first command line argument is called argv[0] and its type is string so you can print it out with a %s format code. You should also recall that the character combination \n means newline. 

Comments:

I could have sworn there was something about an iterator syntax difference between C and Go, I thought it was ++, but it seems to work here. Maybe '+='? 
*/

package main

import (
  "fmt"
  "os"
)

func main() {
  var i, argCount int
  var argString []string  

  // Store the string slice that is returned by the os package variable Args. 
	// Store its length in argCount.  
  argString = os.Args
  argCount = len(argString)

  // For personal formatting.
  fmt.Printf("\nThere were %d arguments read:\n\n", argCount)

  for i < argCount {
    fmt.Printf("Args[%d]: %s\n\n", i, argString[i])

    i++
  }
}
