/*
Problem #29: Find the Average
Author: Corey Prak
Date Created: 04/14/2014

Description: Compute the average of four integer values read from a file. Remember that the average of a number isn't necessarily an integer. 

Comments:
Stuck with only one error variable, thinking about some other way to streamline the error handling. Also extracted logic from being hardcoded in the program, such as the number of values to be read in the program and the file name.

Didn't notice at first that the loop ran 5 times instead of 4. Actually added another value to file and noticed a change, which shouldn't have happened.

There's a format specifier that will leave out trailing, unnecessary 0s I think! Will come back...
*/

package main

import(
  "fmt"
  "os"
  "io"
)

func main() {
  var filePointer *os.File
  var err error
  var i, numberOfValues, sum, numberFromFile int
  var fileName string  

  numberOfValues = 4
  fileName = "testdata26"

  filePointer, err = os.Open(fileName)
  if err != nil {
    fmt.Printf("\n\nOpen Error: %v\n\nExiting...\n\n", err)
    return
  } 

  fmt.Printf("\nContents of %s:\n", fileName)

  for ; i < numberOfValues; i++ {
		_, err = fmt.Fscanf(filePointer, "%d", &numberFromFile)
		if ((err != nil) && (err != io.EOF)) {
			fmt.Printf("\n\nScan Error: %v\n\nExiting...\n\n", err)
			return
		}

    if err != io.EOF {
			fmt.Printf("\n%d", numberFromFile)

			sum += numberFromFile
    }
	}
  
  fmt.Printf("\n\nThe average value of the %d values is %f.\n\n", numberOfValues, (float32(sum) / float32(numberOfValues)))

  err = filePointer.Close()  
}

