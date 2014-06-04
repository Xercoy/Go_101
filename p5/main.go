/*
Problem #5: The Sum of Two Values
Author: Corey Prak
Date Created: 04/03/2014

Description: Use Scanf to read in two integers, store the sum in a 
variable of the same word, print the variable to stdout. 

Comments:
06/03/2014

Added comments, changed code drastically. 
*/

package main

import "fmt"

func main(){
 
//  Declare two variables which are defaulted to a value of 0.  
  var num1, num2, sum int

/*  Request for the user to enter two integer values. As needed by 
    the Scanf function, the next two arguments must be an address to 
    a particular location in memory that is of a matching type. %d
    requires an integer value, and that is what was provided. Also,
    note how Scanf is able to read in more than one integer. The format
    specifiers must be separated by a space when doing this.  */
  fmt.Println("Please enter two integer values:")
  fmt.Scanf("%d %d", &num1, &num2)

  sum = num1 + num2

/*  Use the %d format specifier to ensure that the next argument is 
    an expression that reports out a value with the particular type,
    in this case, an integer.  */
  fmt.Printf("The sum of the two integers is %d.\n", sum)
}
