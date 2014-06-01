/*
Problem #9: Using a For Loop
Author: Corey Prak
Date Created: 04/16/2014
Comments:

*/

package main

import (
  "fmt"
  "os"
)

func main() {

  var fileNum int = 0

  filePtr, openErr := os.Open("testdata9")
  if openErr != nil {
    fmt.Println(openErr)
  }

  _, EOFSubstitute := fmt.Fscanf(filePtr, "%d", &fileNum);
  for  EOFSubstitute == nil {  
    fmt.Printf("%d\n", fileNum)
    _, EOFSubstitute = fmt.Fscanf(filePtr, "%d", &fileNum)
  }
}
