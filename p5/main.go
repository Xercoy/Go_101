/*
Problem #5: The Sum of Two Values
Author: Corey Prak
Date Created: 04/03/2014
Comments:
*/

package main

import "fmt"

func main(){
  var n1 int
  n2 := 0

  n1 = 0

  fmt.Println("Please enter two integer values:")
  fmt.Scanf("%d%d", &n1, &n2)

  fmt.Printf("The sum of the two integers is %d.\n", (n1 + n2))
}