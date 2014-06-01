/*
Problem #12: Positive, Negative, or Zero?
Author: Corey Prak
Date Created: 04/18/2014

Write a C program that reads a single integer number from the keyboard via the scanf function, determines if the number is either, positive, zero, or negative and then prints out a message such as: “The number is positive.” or “The number is zero.” or “The number is negative.”

Comments:
*/

package main

import "fmt"

func main() {
  var userInput int 
  var scanErr error

  fmt.Printf("\nPlease enter an integer value: ")

  _, scanErr = fmt.Scanf("%d", &userInput)
  if scanErr != nil {
    fmt.Printf("\nERROR: %s\n\n", scanErr) 
    return
  }

  if userInput > 0 {
    fmt.Printf("The number is positive.\n\n")
  } else if userInput == 0 {
    fmt.Printf("The number is zero.\n\n")
  } else if userInput < 0 {
    fmt.Printf("The number is negative.\n\n")
  }
}
