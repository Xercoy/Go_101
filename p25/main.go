/*
Problem #25: Sum of a Bunch
Author: Corey Prak
Date Created: 04/03/2014

Description: Create a file and call it testdata25. The file should contain an unknown number of integer values. Write a program that will read in these numbers and output the sum. You should not use an array to do this problem.

Comments:
*/

package main

import (
	//Printing to console
	"fmt"

	//File manipulation
	"os"

	//EOF variable
	"io"
)

func main() {
	var filePointer *os.File
	var err error
	var sum, numberFromFile int

	filePointer, err = os.Open("testdata25")
	if err != nil {
		fmt.Printf("\nOpen Error: %v\n\nExiting...\n\n", err)
		return
	}

	fmt.Printf("\nNumber values from file:\n")

	for err == nil {
		_, err = fmt.Fscanf(filePointer, "%d", &numberFromFile)

		if (err != nil) && (err != io.EOF) {
			fmt.Printf("\n\nScan Error: %v\n\nExiting...", err)
			return
		}

		if err != io.EOF {
			fmt.Printf("\n%d", numberFromFile)

			sum = sum + numberFromFile
		}
	}

	fmt.Printf("\n\nSum of integers contained in testdata25: %d\n\n", sum)

	err = filePointer.Close()
	if err != nil {
		fmt.Printf("\n\nClose Error: %v\n\nExiting...\n\n", err)
		return
	}
}
