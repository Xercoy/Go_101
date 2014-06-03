/*
Problem #4: The scanf Function
Author: Corey Prak
Date Created: 04/03/2014
Comments:

06/02/2014

- Added comments

- Changed the declaration of the userInput variable to be explicit rather than using short declaration. 

- Scan, Scanf, and Scanln read from os.Stdin.

- Scanf returns the number of items successfully scanned. For simplicity, I consider it beyond the scope of this example.
*/

package main

import "fmt"

func main(){
  var userInput String

  fmt.Printf("Enter any sequence of characters:")

/* Here is the first time that we're taking in user input. As you can see, the first argument is a string literal with a format specifier%s (uninterpreted bytes of the string). The next argument is the address of the userInput variable. The Scanf function must be given a location in memory to place the data it received from standard input, this is exactly why an address is needed. Both the necessary type of the format specifier and variable must match or be compatable. */ 
  fmt.Scanf("%s", &userInput)

/* The \n is another special sequence of characters called an escape sequence. Google is your friend! */
  fmt.Printf("User input:%s\n", userInput)
}
