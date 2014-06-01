/*
Problem #17
Author: Corey Prak
Date Created: 05/03/2014

Description: 

Write a program that counts the characters entered into the program from the keyboard until EOF is reached. Use the getchar function. Your program should output the number of characters entered.

Comments:

This took the most time, just becuse I had to learn about the different libraries and how buffers worked. Definitely satisfying to complete, I had so many "Aha!" moments. 
*/

package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  var reader *bufio.Reader
  var readErr error
  var totalByteCount int
  var delimitedGroupofBytes []byte

  reader = bufio.NewReader(os.Stdin)

  fmt.Println("Enter as many characters as you'd like, EOF/Ctrl-D will signal it to stop.")

  for readErr == nil {  
    delimitedGroupofBytes, readErr = reader.ReadBytes('\n')

//    Debugging, just so that useful information could be viewed. 
//    fmt.Println("Delimited group of bytes read: ", delimitedGroupofBytes, ", length: ", len(delimitedGroupofBytes))

    if readErr == nil {
      totalByteCount = totalByteCount + 1
    }
  }

  fmt.Printf("\n\nEOF Reached. A total of %d characters were read.\n\n", totalByteCount)
}
