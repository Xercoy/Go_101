/*
Problem #8: One Horizontal Line of Asterisks
Author: Corey Prak
Date Created: 04/03/2014

Description
Using the emacs editor, create a file called testdata8 that has a single integer value stored into it. The integer will be a number between 1 and 30 inclusive. Create a C program that uses the fscanf function to read in this integer value and then prints a horizontal line containing that many asterisks’ (*). Use a for-loop in your program. 

Comments:

Attempted to cast string from file read to integer. didn't work out, was 
indicated by printf %T format specifier that it was of type []uint8, however
error message upon cast attempt claimed it was of type []byte. I can't 
explain this. 

Decided to dig deeper and found the good ol' Atoi function. 

Dug a little deeper and found that FSCANF EXISTS. WHUUUT
*/

package main

import(
"fmt"
"os"
)

func main(){
  fileNum := 0
  // Iterator for the for loop.
  i := 0;

  // Open the file, return address to filePtr, fErr if possible.
	filePtr, fErr := os.Open("testdata8")
  
  if fErr != nil {
		fmt.Println(fErr)
  }

  _, scanErr := fmt.Fscanf(filePtr, "%d", &fileNum)
  if scanErr != nil {
    fmt.Println(scanErr)
  }
  

  for i = 0; i < fileNum  ; i = i+1 {
	  fmt.Printf("*")
  }
  
		fmt.Printf("\n")
}
