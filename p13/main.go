/*
Problem #13: The abs Function
Author: Corey Prak
Date Created: 04/19/2014

Description:
Write a C program that reads a single integer value from the keyboard. The number could be either a negative number, a positive number, or zero. Your program should call the absolute value function – abs – and it then prints out the absolute value of number entered. You can browse appendix D to determine which header file you should include. Your textbook is your best friend.  I do not want you to create your own absolute value function. Make sure that your code looks snappy. What was the trouble in River City?

Comments: I thought it was pretty cool that I finished this quickly, but at the same time i wish that I really didn't have to rely on the %q format specifier when I can't even tell anyone how it works yet =/.
*/

package main

import(
 "fmt"
 "math"
) 

func main(){
  // Since we're going to use the abs function which takes in a float64 int, that's texactly the type of data i'm going to supply it with.
  var userInput float64
  var scanErr error

  fmt.Printf("\nPlease enter an integer value: ")
  _, scanErr = fmt.Scanf("%f", &userInput)
  if scanErr != nil {
    fmt.Printf("\nERROR: %q.\n\nExiting...\n\n", scanErr)
    return
  }

  fmt.Printf("\nThe absolute value of %f is %f.\n\n", userInput, math.Abs(userInput)) 
}
