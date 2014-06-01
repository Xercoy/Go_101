/*
Problem #4: The scanf Function
Author: Corey Prak
Date Created: 04/03/2014
Comments:
*/

package main

import "fmt"

func main(){
  userInput := ""

  fmt.Printf("Enter some kind of input:")
  fmt.Scanf("%s", &userInput)
  fmt.Printf("User input:%s\n", userInput)
}