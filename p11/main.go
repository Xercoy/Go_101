/*
Problem #11: Equal to Zero?
Author: Corey Prak
Date Created: 04/17/2014
Comments:
*/

package main

import "fmt"

func main() {
  var userInput int
  var successfulScans int
  var scanErr error

  fmt.Printf("\nPlease enter a single integer value: ")

  successfulScans, scanErr = fmt.Scanf("%d", &userInput)
  if scanErr != nil {
    fmt.Println("ERROR: ", scanErr)
  }

  if successfulScans == 0 {
    fmt.Printf("Invalid input entered. Exiting.\n\n")
    return
  }

  if userInput == 0 {
    fmt.Printf("The number is equal to zero.\n\n")
  } else {
    fmt.Printf("The number is not equal to zero.\n\n")
	}
}
