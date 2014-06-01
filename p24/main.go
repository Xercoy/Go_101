/*
Problem #24: One Dimensional Array
Author: Corey Prak
Date Created: 04/03/2014

Description: Create a file and call it testdata24. The file should contain exactly 15 integer values. Read these 15 numbers into a one-dimensional array. Then print the numbers out in the reverse order from which they were entered.
Use an array. I do not want you to create 15 distinct variables. Read your book. Become independent.

Comments:

Attempted to reuse i in the second for loop without assigning it to an appropriate value. I expected it to be 14 after the first loop was done, however I totally did not account for the last increment which would make i 15, satisfying the condition for the loop to end. Thus, in the second for loop I explicitly assigned i to be one less than the current value at that point. 

Might be ugly, but this problem could also have been done using only one for loop. Before the end of the first loop, add in something like 'fmt.Printf("numberArray[%d]: %d", 14 - i, numberArray[14 - i])'. Yes, hard coding logic = bad, but you could use len(). If I understand correctly, len() returns the size of its argument in bytes. Maybe len wouldn't work if we were dealing with a data type that wasn't so convenient, like strings. Would len() of a int32 array correctly return the number of elements, or the size in bytes that it takes up?   
*/

package main

import (
  "fmt"
  "os"
)

func main() {
  var numberArray [15]int
  var openErr, scanErr, closeErr error
  var filePointer *os.File
  var i int

  filePointer, openErr = os.Open("testdata24")
  if openErr != nil {
    fmt.Printf("\nOpen Error: %v\n\nExiting...\n\n", openErr)
    return
  }

  fmt.Printf("\nLength of numberArray, which will be used to store our 15 integers from the file testdata24: %d\n", len(numberArray))

  for i = 0; i < 15; i++ {
    _, scanErr = fmt.Fscanf(filePointer, "%d", &numberArray[i])

    if scanErr != nil {
      fmt.Printf("\nScan Error: %v\n\nExiting...\n\n", scanErr)
      return
    } 
  } 

  //Reuse i!!! Why not?
  for i -= 1; i >= 0; i-- {
    fmt.Printf("\nnumberArray[%d]: %d", i, numberArray[i])
  }

  closeErr = filePointer.Close()
  if closeErr != nil {
    fmt.Printf("\nClose Error: %V\n\nExiting...\n\n", closeErr)
  }

  //Formating...
  fmt.Printf("\n\n")  
}
