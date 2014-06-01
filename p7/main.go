/*
Problem #7: Bigger than 100?
Author: Corey Prak
Date Created: 04/13/2014

Description:
Write a C program that reads a number from the keyboard via the scanf function, and then prints the message The number is bigger than 100 if it is. If the number is 100 or less than 100 then you should print out the message: The number is not bigger than 100. Make certain that your code looks good. No ragged indents.

Comments:
MUST REMEMBER THAT THE PACKAGE NAME OF AN IMPORT MUST BE A STRING.

Wth, compiler will kick your ass if the else if isn't on the same line as 
the previous block's closing bracket.
*/

// Import main package.
package main

// import fmt package for I/O.
import "fmt"

func main(){
  var userInput int32 = 0;

  // Ask for integer input, store value in userInput.
  fmt.Println("Please enter an integer value:")
  fmt.Scanf("%d", &userInput)

  // Is userInput larger than 100? Indicate whether it is or isn't so.
  if userInput > 100 {
  	fmt.Println("The number is bigger than 100.")
  } else if userInput <= 100 {
    fmt.Println("The number is not bigger than 100.")
  }
}
