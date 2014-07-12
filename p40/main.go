/*******************************************************************************
Problem #40: Greated Common Divisor (Brute Force Method)
Author: Corey Prak
Date Created: 06/12/2014
Comments:
- GCD function made use of the variable lesserNumber to keep things DRY.

- The chunk of if statements in the main function are really ugly...
*******************************************************************************/

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func GCD(num1 int, num2 int) int {
	// Make an initial assignment
	lesserNumber := num1

	// Swap if necessary
	if num2 < num1 {
		lesserNumber = num2
	}

	for i := lesserNumber; i > 1; i-- {
		if (num1%i == 0) && (num2%i == 0) {
			return i
		}
	}

	// Bothered me not to have a return statement, so I took = out of
	// the for loop
	return 1
}

func main() {
	var filePtr *os.File
	var firstNumber, secondNumber int
	var openErr, firstNumberScanErr, secondNumberScanErr error

	filePtr, openErr = os.Open(os.Args[1])
	if openErr != nil {
		log.Fatalln(openErr)
	}
	defer filePtr.Close()

	for true {
		_, firstNumberScanErr = fmt.Fscanf(filePtr, "%d", &firstNumber)
		_, secondNumberScanErr = fmt.Fscanf(filePtr, "%d", &secondNumber)

		//Guaranteed pairs, so firstNumberScanErr will always be EOF when
		//reading is done.
		if firstNumberScanErr == io.EOF {
			break

			// If there is an error reading the first number of the pair, print error
		} else if firstNumberScanErr != nil {

			// If there is an error with both numbers, print both their errors
			// else print only the first one.
			if secondNumberScanErr != nil {
				log.Fatalln(firstNumberScanErr, "\n", secondNumberScanErr)
			} else {
				log.Fatalln(firstNumberScanErr)
			}

			// If there is an error reading the second number of the pair, print error
		} else if secondNumberScanErr != nil {
			log.Fatalln(firstNumberScanErr)
		}

		fmt.Printf("\nThe GCD of %d and %d is %d\n", firstNumber, secondNumber, GCD(firstNumber, secondNumber))
	}

	fmt.Printf("\n")
}
