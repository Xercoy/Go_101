/*
Problem #22: Reverse Command Line
Author: Corey Prak
Date Created: 05/10/2014
Comments:
*/

package main

import (
  "fmt"
  "os"
)

func main() {
  var argStringSlice []string
  var argCount, i int

  argStringSlice = os.Args
  argCount = len(argStringSlice)

  fmt.Printf("\n Number of command line arguments: %d\n\n", argCount)

  for i = argCount - 1; i >= 0; i = i - 1 {
    fmt.Printf("Argv[%d]: %s\n\n", i, argStringSlice[i])
  }  
}
