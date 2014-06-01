/*
Problem #10: Sum of Twenty 
Author: Corey Prak
Date Created: 04/17/2014

Description: Problem 10: Sum of Twenty

Write a C program that reads in 20 positive integer values from a file called testdata10 and computes their sum, and then prints the value of the sum onto the screen.  Your program should not use 20 distinct variables. Your program should not use an array.

Comments:
MUST READ INTO PACKAGES. I hate short declarations. They're convenient, not useful. This was the first time that I was able to actually declare a file pointer normally. Will write about it...
*/

package main

import(
  "fmt"
  "os"
)

func main(){
  var filePointer * os.File
  var fileNumber, sum int
  var openErr, countErr error

  filePointer, openErr = os.Open("testdata10")
  if openErr != nil {
    fmt.Println(openErr)
  } 

  _, countErr = fmt.Fscanf(filePointer, "%d", &fileNumber)
  if countErr != nil {
    fmt.Println(countErr)
    return 
  }

  for countErr == nil {
    sum += fileNumber
		_, countErr = fmt.Fscanf(filePointer, "%d", &fileNumber)
  }

  fmt.Printf("The sum of the 20 integers read from testdata10 is %d\n\n", sum)
}
