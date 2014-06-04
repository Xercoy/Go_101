/*******************************************************************************
Problem #7: Bigger than 100?
Author: Corey Prak
Date Created: 04/13/2014

Description:
Write a program that reads a number from the keyboard via the scanf function, 
and then prints the message The number is bigger than 100 if it is. If the 
number is 100 or less than 100 then you should print out the message: The 
number is not bigger than 100. Make certain that your code looks good. 
No ragged indents.

Comments:
06/04/2014

Added descriptive and helpful comments.

Made some code changes. 
*******************************************************************************/

// Import main package.
package main

// import fmt package for I/O.
import "fmt"

func main(){
  var userInput int32;

  // Ask for integer input, store value in userInput.
  fmt.Println("Please enter an integer value:")
  fmt.Scanf("%d", &userInput)

  /* If the expression between an if and { bracket is evaluated to be true, 
     then statements in its code block will be executed. If the condition is
	   not true, then successive 'else if' conditions will be evaluated 
     and its respective code block will run if true. If there is an else 
     condition that is 'hit' and no previous conditions have been true, 
     the code of the else block will run. 

     In our case, the first condition verifies whether userInput is greater
     than 100, if so, the code in its block will run. If that isn't true, 
     then the expression of the else if block will be evalulated. If the 
     value of userInput is EQUAL TO OR LESS THAN 100, then*/
  if userInput > 100 {
  	fmt.Println("The number is bigger than 100.")
  } else if userInput <= 100 {
    fmt.Println("The number is not bigger than 100.")
  }
}
