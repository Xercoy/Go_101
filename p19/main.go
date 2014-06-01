/*
Problem #19: Area of a Rectangle
Author: Corey Prak
Date Created: 04/21/2014
Comments:
*/

package main

import "fmt"

func main(){

  var length, width float32
  var scanErr error

  fmt.Printf("\nPlease enter two floating point numbers, length first, then width:\n")

  //Format string must have format specifiers separated with a space.
  _, scanErr = fmt.Scanf("%f %f", &length, &width)
  if scanErr != nil {
    fmt.Printf("\nERROR:%q\n\nExiting...\n\n", scanErr)
    return
  }

  //No need for an area variable, include the expression which will report out a value to the format specifier.
  fmt.Printf("\nThe area of a retangle with a length of %f and width of %f is %f.\n\n", length, width, length * width)
}
