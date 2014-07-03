/*******************************************************************************
 Problem #36: Brute Force Primes
Author: Corey Prak
Date Created: 06/30/2014
 Date Completed: 07/03/2014
Comments:
- Avoided short declarations for some things, I forget the types of so many
things already.

- Used a range since I never did before. Surprising!

- I wanted to DRY error checking but not sure if this was a best practice
since we are allowing another function other than main to quit. Might be
confusing at a larger scale.
*******************************************************************************/

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func IsPrime(number int) bool {
	var numberOfDivisors int

	for i := 1; i <= number; i++ {
		if numberOfDivisors > 2 {
			return false
		} else if (number % i) == 0 {
			numberOfDivisors++
		}
	}

	return true
}

func ErrorHandler(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	var file *os.File
	var openErr, scanErr, closeErr error
	var intSlice []int

	if len(os.Args) == 1 {
		log.Fatalln("ERROR, file to be opened must be provided.")
	}

	file, openErr = os.Open(os.Args[1])
	ErrorHandler(openErr)

	for i := 0; ; {
		_, scanErr = fmt.Fscanf(file, "%d", &i)
		if scanErr == io.EOF {
			break
		}

		intSlice = append(intSlice, i)
	}

	fmt.Printf("\nIntegers read from file: %v\n", intSlice)

	for _, value := range intSlice {
		if value == 0 {
			fmt.Println("\n 0 is neither prime nor composite.")
		} else if IsPrime(value) {
			fmt.Printf("\n%d is prime.", value)
		} else {
			fmt.Printf("\n%d is not prime.", value)
		}

	}
	fmt.Printf("\n\n")

	closeErr = file.Close()
	ErrorHandler(closeErr)
}
